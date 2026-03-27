package services

import (
	"Synconomics/models"
	"Synconomics/repositories"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type aiService struct {
	repo            repositories.AIRepository
	transactionRepo repositories.TransactionRepository
	expenseRepo     repositories.ExpenseRepository
	businessRepo    repositories.BusinessRepository
	productRepo     repositories.ProductRepository
}

func NewAIService(
	repo repositories.AIRepository,
	transactionRepo repositories.TransactionRepository,
	expenseRepo repositories.ExpenseRepository,
	businessRepo repositories.BusinessRepository,
	productRepo repositories.ProductRepository,
) AIService {
	return &aiService{
		repo:            repo,
		transactionRepo: transactionRepo,
		expenseRepo:     expenseRepo,
		businessRepo:    businessRepo,
		productRepo:     productRepo,
	}
}

func (s *aiService) CreateSession(userID, businessID uint, sessionType string) (*models.AISession, error) {
	session := &models.AISession{
		UserID:     userID,
		BusinessID: businessID,
		Type:       models.AISessionType(sessionType),
	}
	if err := s.repo.CreateSession(session); err != nil {
		return nil, err
	}
	return session, nil
}

func (s *aiService) getTools() []*genai.Tool {
	return []*genai.Tool{
		{
			FunctionDeclarations: []*genai.FunctionDeclaration{
				{
					Name:        "get_user_and_business_info",
					Description: "Mendapatkan profil user (nama, email) dan daftar bisnis yang ia miliki (termasuk ID bisnis). Panggil ini terlebih dahulu jika Anda belum tahu ID bisnis user.",
				},
				{
					Name:        "get_market_trends",
					Description: "Mendapatkan daftar 10 tren pasar teratas saat ini untuk memberikan konteks eksternal.",
				},
				{
					Name:        "get_business_products",
					Description: "Mendapatkan daftar lengkap produk dari sebuah bisnis. Membutuhkan ID bisnis yang valid.",
					Parameters: &genai.Schema{
						Type: genai.TypeObject,
						Properties: map[string]*genai.Schema{
							"business_id": {
								Type:        genai.TypeInteger,
								Description: "ID unik bisnis (bukan ID user). Dapatkan dari get_user_and_business_info.",
							},
						},
						Required: []string{"business_id"},
					},
				},
				{
					Name:        "get_business_transactions",
					Description: "Mendapatkan seluruh riwayat transaksi penjualan (POS) sebuah bisnis. Gunakan ini untuk menganalisis pendapatan atau performa penjualan.",
					Parameters: &genai.Schema{
						Type: genai.TypeObject,
						Properties: map[string]*genai.Schema{
							"business_id": {
								Type:        genai.TypeInteger,
								Description: "ID unik bisnis.",
							},
						},
						Required: []string{"business_id"},
					},
				},
				{
					Name:        "update_user_profile",
					Description: "Memperbarui profil user (nama, email). Kosongkan jika tidak ingin diubah.",
					Parameters: &genai.Schema{
						Type: genai.TypeObject,
						Properties: map[string]*genai.Schema{
							"name":  {Type: genai.TypeString, Description: "Nama baru user."},
							"email": {Type: genai.TypeString, Description: "Email baru user."},
						},
					},
				},
				{
					Name:        "update_business_info",
					Description: "Memperbarui informasi bisnis. Membutuhkan business_id.",
					Parameters: &genai.Schema{
						Type: genai.TypeObject,
						Properties: map[string]*genai.Schema{
							"business_id": {Type: genai.TypeInteger, Description: "ID bisnis yang akan diupdate."},
							"name":        {Type: genai.TypeString, Description: "Nama baru bisnis."},
							"description": {Type: genai.TypeString, Description: "Deskripsi baru."},
							"address":     {Type: genai.TypeString, Description: "Alamat baru."},
							"category":    {Type: genai.TypeString, Description: "Kategori baru."},
						},
						Required: []string{"business_id"},
					},
				},
				{
					Name:        "update_product_info",
					Description: "Memperbarui informasi produk. Membutuhkan product_id.",
					Parameters: &genai.Schema{
						Type: genai.TypeObject,
						Properties: map[string]*genai.Schema{
							"product_id":  {Type: genai.TypeInteger, Description: "ID produk yang akan diupdate."},
							"name":        {Type: genai.TypeString, Description: "Nama baru produk."},
							"description": {Type: genai.TypeString, Description: "Deskripsi baru."},
							"price":       {Type: genai.TypeNumber, Description: "Harga baru."},
							"stock":       {Type: genai.TypeInteger, Description: "Jumlah stok baru."},
							"min_stock":   {Type: genai.TypeInteger, Description: "Minimum stok baru."},
						},
						Required: []string{"product_id"},
					},
				},
			},
		},
	}
}

func (s *aiService) callInternalAPI(ctx context.Context, method string, path string, body interface{}, token string) (interface{}, error) {
	baseURL := os.Getenv("API_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080/api"
	}

	var bodyReader io.Reader
	if body != nil {
		jsonBody, _ := json.Marshal(body)
		bodyReader = strings.NewReader(string(jsonBody))
	}

	fmt.Printf("DEBUG: AI API Call: [%s] %s%s\n", method, baseURL, path)

	req, err := http.NewRequestWithContext(ctx, method, baseURL+path, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("DEBUG: AI API Error: %v\n", err)
		return map[string]interface{}{"error": "connection failed", "details": err.Error()}, nil
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{"error": "read failed"}, nil
	}

	if resp.StatusCode >= 400 {
		fmt.Printf("DEBUG: AI API Non-OK Status: %d, Body: %s\n", resp.StatusCode, string(respBody))
		var errData map[string]interface{}
		json.Unmarshal(respBody, &errData)
		return map[string]interface{}{"error": "request failed", "status": resp.StatusCode, "data": errData}, nil
	}

	var result interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return map[string]interface{}{"error": "parse failed", "raw": string(respBody)}, nil
	}

	return result, nil
}

// partJSON tetap sama
type partJSON struct {
	Type     string                 `json:"type"`
	Text     string                 `json:"text,omitempty"`
	FuncName string                 `json:"func_name,omitempty"`
	Args     map[string]interface{} `json:"args,omitempty"`
	Response map[string]interface{} `json:"response,omitempty"`
}

func (s *aiService) saveParts(sessionID uint, role string, parts []genai.Part) error {
	var pj []partJSON
	var plainText string
	isSimple := true

	for _, p := range parts {
		item := partJSON{}
		switch v := p.(type) {
		case genai.Text:
			item.Type = "text"
			item.Text = string(v)
			plainText += string(v)
		case genai.FunctionCall:
			item.Type = "func_call"
			item.FuncName = v.Name
			item.Args = v.Args
			isSimple = false
		case genai.FunctionResponse:
			item.Type = "func_resp"
			item.FuncName = v.Name
			item.Response = v.Response
			isSimple = false
		default:
			item.Type = "unknown"
			isSimple = false
		}
		pj = append(pj, item)
	}

	content := plainText
	if !isSimple {
		bytes, _ := json.Marshal(pj)
		content = "JSON_PARTS:" + string(bytes)
	}

	msg := &models.AIMessage{
		AISessionID: sessionID,
		Role:        role,
		Content:     content,
	}
	return s.repo.SaveMessage(msg)
}

func (s *aiService) loadHistory(sessionID uint) ([]*genai.Content, error) {
	storedMessages, err := s.repo.GetMessagesBySessionID(sessionID)
	if err != nil {
		return nil, err
	}

	var history []*genai.Content
	for _, msg := range storedMessages {
		var parts []genai.Part
		if strings.HasPrefix(msg.Content, "JSON_PARTS:") {
			jsonStr := strings.TrimPrefix(msg.Content, "JSON_PARTS:")
			var pj []partJSON
			if err := json.Unmarshal([]byte(jsonStr), &pj); err == nil {
				for _, item := range pj {
					switch item.Type {
					case "text":
						parts = append(parts, genai.Text(item.Text))
					case "func_call":
						parts = append(parts, genai.FunctionCall{Name: item.FuncName, Args: item.Args})
					case "func_resp":
						parts = append(parts, genai.FunctionResponse{Name: item.FuncName, Response: item.Response})
					}
				}
			}
		}

		if len(parts) == 0 {
			parts = append(parts, genai.Text(msg.Content))
		}

		history = append(history, &genai.Content{
			Role:  msg.Role,
			Parts: parts,
		})
	}
	return history, nil
}

func (s *aiService) Chat(sessionID uint, userMessage string, token string) (*models.AIMessage, error) {
	session, err := s.repo.GetSessionByID(sessionID)
	if err != nil {
		return nil, err
	}

	// Save user message
	if err := s.saveParts(sessionID, "user", []genai.Part{genai.Text(userMessage)}); err != nil {
		return nil, err
	}

	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.5-flash") // Defaulting back to stable 1.5-flash for tools
	model.Tools = s.getTools()

	var systemInstruction string
	switch session.Type {
	case models.IdeaGeneration:
		systemInstruction = "Anda adalah konsultan bisnis AI ahli dalam Idea Generation. Anda dapat mengakses data real-time melalui tools yang tersedia."
	case models.Validation:
		systemInstruction = "Anda adalah konsultan bisnis AI ahli dalam Business Validation. Gunakan data bisnis untuk memvalidasi ide."
	case models.Strategy:
		systemInstruction = "Anda adalah konsultan bisnis AI ahli dalam Strategi Bisnis. Analisis data transaksi dan pengeluaran untuk memberikan strategi."
	default:
		systemInstruction = "Anda adalah asisten AI bisnis yang cerdas. Selalu gunakan tools untuk memberikan informasi yang akurat tentang data bisnis user."
	}
	model.SystemInstruction = &genai.Content{Parts: []genai.Part{genai.Text(systemInstruction)}}

	cs := model.StartChat()
	history, _ := s.loadHistory(sessionID)
	if len(history) > 0 && history[len(history)-1].Role == "user" {
		cs.History = history[:len(history)-1]
	} else {
		cs.History = history
	}

	resp, err := cs.SendMessage(ctx, genai.Text(userMessage))
	if err != nil {
		return nil, fmt.Errorf("gemini sdk error: %v", err)
	}

	for {
		if len(resp.Candidates) == 0 {
			return nil, fmt.Errorf("no candidates returned")
		}

		modelContent := resp.Candidates[0].Content
		var funcResponses []genai.Part
		hasFuncCall := false

		for _, part := range modelContent.Parts {
			if funcCall, ok := part.(genai.FunctionCall); ok {
				hasFuncCall = true
				var result interface{}
				fmt.Printf("DEBUG: AI Calling %s(%v)\n", funcCall.Name, funcCall.Args)

				switch funcCall.Name {
				case "get_user_and_business_info":
					result, _ = s.callInternalAPI(ctx, "GET", "/auth/me-with-businesses", nil, token)
				case "get_market_trends":
					result, _ = s.callInternalAPI(ctx, "GET", "/market-trends/top", nil, token)
				case "get_business_products", "get_business_transactions", "get_business_expenses":
					bizIDRaw, ok := funcCall.Args["business_id"]
					if !ok {
						result = map[string]interface{}{"error": "missing business_id parameter"}
					} else {
						bizID := strings.TrimSuffix(fmt.Sprintf("%v", bizIDRaw), ".0")
						path := ""
						switch funcCall.Name {
						case "get_business_products":
							path = "/products/business/" + bizID
						case "get_business_transactions":
							path = "/transactions/business/" + bizID
						case "get_business_expenses":
							path = "/expenses/business/" + bizID
						}
						result, _ = s.callInternalAPI(ctx, "GET", path, nil, token)
					}
				case "update_user_profile":
					result, _ = s.callInternalAPI(ctx, "PUT", "/auth/profile", funcCall.Args, token)
				case "update_business_info":
					bizID := strings.TrimSuffix(fmt.Sprintf("%v", funcCall.Args["business_id"]), ".0")
					result, _ = s.callInternalAPI(ctx, "PUT", "/business/"+bizID, funcCall.Args, token)
				case "update_product_info":
					prodID := strings.TrimSuffix(fmt.Sprintf("%v", funcCall.Args["product_id"]), ".0")
					result, _ = s.callInternalAPI(ctx, "PUT", "/products/"+prodID, funcCall.Args, token)
				}

				funcResponses = append(funcResponses, genai.FunctionResponse{
					Name:     funcCall.Name,
					Response: map[string]interface{}{"result": result},
				})
			}
		}

		if !hasFuncCall {
			if err := s.saveParts(sessionID, "model", modelContent.Parts); err != nil {
				return nil, err
			}

			var modelResponseText string
			for _, p := range modelContent.Parts {
				if textPart, ok := p.(genai.Text); ok {
					modelResponseText += string(textPart)
				}
			}

			return &models.AIMessage{
				AISessionID: sessionID,
				Role:        "model",
				Content:     modelResponseText,
			}, nil
		}

		if err := s.saveParts(sessionID, "model", modelContent.Parts); err != nil {
			return nil, err
		}
		if err := s.saveParts(sessionID, "user", funcResponses); err != nil {
			return nil, err
		}

		resp, err = cs.SendMessage(ctx, funcResponses...)
		if err != nil {
			return nil, fmt.Errorf("gemini sdk error (after tool): %v", err)
		}
	}
}

func (s *aiService) FinalizeSessionResult(sessionID uint) (*models.AIResult, error) {
	messages, err := s.repo.GetMessagesBySessionID(sessionID)
	if err != nil || len(messages) == 0 {
		return nil, fmt.Errorf("sesi kosong atau tidak ditemukan")
	}

	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.5-flash")
	model.ResponseMIMEType = "application/json"

	var fullConversation string
	for _, m := range messages {
		content := m.Content
		if strings.HasPrefix(content, "JSON_PARTS:") {
			var pj []partJSON
			if err := json.Unmarshal([]byte(strings.TrimPrefix(content, "JSON_PARTS:")), &pj); err == nil {
				content = ""
				for _, item := range pj {
					if item.Type == "text" {
						content += item.Text
					}
				}
			}
		}
		fullConversation += fmt.Sprintf("[%s]: %s\n", m.Role, content)
	}

	prompt := fmt.Sprintf("Analyze the conversation and provide a JSON result. Conversation:\n%s", fullConversation)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}

	var jsonOutput string
	for _, part := range resp.Candidates[0].Content.Parts {
		if textPart, ok := part.(genai.Text); ok {
			jsonOutput += string(textPart)
		}
	}

	result := &models.AIResult{
		AISessionID: sessionID,
		Data:        jsonOutput,
	}
	
	session, _ := s.repo.GetSessionByID(sessionID)
	if session != nil {
		result.ResultType = string(session.Type)
	}

	if err := s.repo.SaveResult(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *aiService) GetSessionMessages(sessionID uint) ([]models.AIMessage, error) {
	return s.repo.GetMessagesBySessionID(sessionID)
}

func (s *aiService) GetSessionResult(sessionID uint) (*models.AIResult, error) {
	return s.repo.GetResultBySessionID(sessionID)
}

func (s *aiService) ChatByRole(userID, businessID uint, sessionType string, message string, token string) (*models.AIMessage, error) {
	session, err := s.repo.GetLatestSession(userID, businessID, models.AISessionType(sessionType))
	if err != nil {
		session, err = s.CreateSession(userID, businessID, sessionType)
		if err != nil {
			return nil, err
		}
	}

	return s.Chat(session.ID, message, token)
}
func (s *aiService) AnalyzeMarketTrends(keywords []string) ([]models.MarketTrend, error) {
	if len(keywords) == 0 {
		return []models.MarketTrend{}, nil
	}

	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.5-flash")
	model.ResponseMIMEType = "application/json"
	
	// Structured prompt for trend analysis
	prompt := fmt.Sprintf(`Analyze the following search keywords collected from various users. 
	1. Normalize them into official product names.
	2. Group similar items.
	3. Assign a 'search_count' (estimate relative popularity based on input frequency).
	4. Assign a 'demand_score' (0.0 to 10.0) based on current global/local trends.
	5. Identify a 'location' if possible, otherwise use 'Global' or 'Indonesia'.
	
	Keywords: %s
	
	Return a JSON array of objects with keys: "product_name", "search_count", "demand_score", "location".`, strings.Join(keywords, ", "))

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}

	var jsonOutput string
	for _, part := range resp.Candidates[0].Content.Parts {
		if textPart, ok := part.(genai.Text); ok {
			jsonOutput += string(textPart)
		}
	}

	var trends []models.MarketTrend
	if err := json.Unmarshal([]byte(jsonOutput), &trends); err != nil {
		return nil, fmt.Errorf("failed to parse AI response: %v, raw: %s", err, jsonOutput)
	}

	return trends, nil
}

func (s *aiService) AuditBusinessReport(userID, businessID uint, token string) (string, error) {
	business, err := s.businessRepo.FindByID(businessID)
	if err != nil {
		return "", fmt.Errorf("failed to find business: %v", err)
	}

	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -30)

	transactions, err := s.transactionRepo.FindByBusinessIDAndDateRange(businessID, startDate, endDate)
	if err != nil {
		return "", fmt.Errorf("failed to fetch transactions: %v", err)
	}

	expenses, err := s.expenseRepo.FindByBusinessIDAndDateRange(businessID, startDate, endDate)
	if err != nil {
		return "", fmt.Errorf("failed to fetch expenses: %v", err)
	}

	products, err := s.productRepo.FindByBusinessID(businessID)
	if err != nil {
		return "", fmt.Errorf("failed to fetch products: %v", err)
	}

	// Prepare data for AI
	data := map[string]interface{}{
		"business":     business,
		"transactions": transactions,
		"expenses":     expenses,
		"products":     products,
		"period":       "last 30 days",
	}

	jsonData, _ := json.MarshalIndent(data, "", "  ")

	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.5-flash")

	systemInstruction := `Role: You are a Senior Financial Business Consultant and Auditor. Your expertise lies in analyzing Point of Sale (POS) data and financial statements to assess a business's health and provide high-impact strategic recommendations.

Objective: Analyze the provided 30-day financial report data and generate a comprehensive Financial Audit Result. You must evaluate the business across four key pillars and provide actionable strategic advice.

Analysis Framework:
You must calculate and interpret the following metrics based on the data provided:
1. Liquidity: Current Ratio, Quick Ratio.
2. Solvency: Debt to Equity Ratio.
3. Profitability: Gross Profit Margin, Net Profit Margin.
4. Efficiency: Inventory Turnover, Asset Turnover.

Strategic Recommendations:
Based on the audit, provide a "Strategic Action Plan" consisting of:
- Cost Reduction Strategy: Identify unnecessary expenditures and provide specific ways to optimize operational costs (OPEX).
- Pricing Optimization: Suggest adjustments to pricing models based on profit margins and sales velocity.
- Investment Priority: Recommend where to allocate capital to maximize business growth.

Tone and Style: Professional, analytical, objective, and constructive. Use clear headings and bullet points for readability. If certain data points are missing, provide an estimate or flag what is needed.`

	model.SystemInstruction = &genai.Content{Parts: []genai.Part{genai.Text(systemInstruction)}}

	prompt := fmt.Sprintf("Analyze the following business data for the last 30 days and provide the audit result:\n\n%s", string(jsonData))

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	var resultText string
	for _, part := range resp.Candidates[0].Content.Parts {
		if textPart, ok := part.(genai.Text); ok {
			resultText += string(textPart)
		}
	}

	return resultText, nil
}
