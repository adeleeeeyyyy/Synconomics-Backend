package services

import (
	"Synconomics/models"

	"github.com/markbates/goth"
)

type AuthService interface {
	Register(name, email, password string) (*models.User, string, error)
	Login(email, password string) (*models.User, string, error)
	HandleGoogleCallback(googleUser goth.User) (*models.User, string, error)
	GetProfile(userID uint) (*models.User, string, error)
}

type ProductService interface {
	CreateProduct(product *models.Product) error
	GetAllProducts() ([]models.Product, error)
	GetProductById(id uint) (*models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error 
}

type BusinessService interface {
	CreateBusiness(business *models.Business) error
	GetAllBusinesses() ([]models.Business, error)
	GetBusinessById(id uint) (*models.Business, error)
	UpdateBusiness(business *models.Business) error
	DeleteBusiness(id uint) error 
}

type TransactionService interface {
	CreateTransaction(transaction *models.Transaction) error
	GetAllTransactions() ([]models.Transaction, error)
	GetTransactionById(id uint) (*models.Transaction, error)
	UpdateTransaction(transaction *models.Transaction) error
	DeleteTransaction(id uint) error 
}

type TransactionItemService interface {
	CreateTransactionItem(item *models.TransactionItem) error
	GetAllTransactionItems() ([]models.TransactionItem, error)
	GetTransactionItemById(id uint) (*models.TransactionItem, error)
	GetTransactionItemsByTransactionId(transactionID uint) ([]models.TransactionItem, error)
	UpdateTransactionItem(item *models.TransactionItem) error
	DeleteTransactionItem(id uint) error 
}

type ExpenseService interface {
	CreateExpense(expense *models.Expense) error
	GetAllExpenses() ([]models.Expense, error)
	GetExpenseById(id uint) (*models.Expense, error)
	GetExpensesByBusinessId(businessID uint) ([]models.Expense, error)
	UpdateExpense(expense *models.Expense) error
	DeleteExpense(id uint) error 
}

type AIService interface {
	CreateSession(userID, businessID uint, sessionType string) (*models.AISession, error)
	Chat(sessionID uint, userMessage string) (*models.AIMessage, error)
	FinalizeSessionResult(sessionID uint) (*models.AIResult, error)
	GetSessionMessages(sessionID uint) ([]models.AIMessage, error)
	GetSessionResult(sessionID uint) (*models.AIResult, error)
	ChatByRole(userID, businessID uint, sessionType string, message string) (*models.AIMessage, error)
}

type SupplyRequestService interface {
	CreateSupplyRequest(req *models.SupplyRequest) error
	GetAllSupplyRequests() ([]models.SupplyRequest, error)
	GetSupplyRequestById(id uint) (*models.SupplyRequest, error)
	GetSupplyRequestsByBusinessId(businessID uint) ([]models.SupplyRequest, error)
	UpdateSupplyRequest(req *models.SupplyRequest) error
	DeleteSupplyRequest(id uint) error
}

type SupplyOfferService interface {
	CreateSupplyOffer(offer *models.SupplyOffer) error
	GetAllSupplyOffers() ([]models.SupplyOffer, error)
	GetSupplyOfferById(id uint) (*models.SupplyOffer, error)
	GetSupplyOffersByBusinessId(businessID uint) ([]models.SupplyOffer, error)
	UpdateSupplyOffer(offer *models.SupplyOffer) error
	DeleteSupplyOffer(id uint) error
}

type SupplyMatchService interface {
	CreateSupplyMatch(match *models.SupplyMatch) error
	GetAllSupplyMatches() ([]models.SupplyMatch, error)
	GetSupplyMatchById(id uint) (*models.SupplyMatch, error)
	GetSupplyMatchesByRequestId(requestID uint) ([]models.SupplyMatch, error)
	GetSupplyMatchesByOfferId(offerID uint) ([]models.SupplyMatch, error)
	UpdateSupplyMatch(match *models.SupplyMatch) error
	DeleteSupplyMatch(id uint) error
}

type ThreadService interface {
	CreateThread(thread *models.Thread) error
	GetAllThreads() ([]models.Thread, error)
	GetThreadById(id uint) (*models.Thread, error)
	UpdateThread(thread *models.Thread) error
	DeleteThread(id uint) error
}

type ReplyService interface {
	CreateReply(reply *models.Reply) error
	GetAllReplies() ([]models.Reply, error)
	GetReplyById(id uint) (*models.Reply, error)
	GetRepliesByThreadId(threadID uint) ([]models.Reply, error)
	UpdateReply(reply *models.Reply) error
	DeleteReply(id uint) error
}

type ProductSearchLogService interface {
	CreateLog(log *models.ProductSearchLog) error
	GetAllLogs() ([]models.ProductSearchLog, error)
	GetLogById(id uint) (*models.ProductSearchLog, error)
	GetLogsByUserId(userID uint) ([]models.ProductSearchLog, error)
}