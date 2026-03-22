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
}

type SupplyRequestService interface {
	CreateSupplyRequest(req *models.SupplyRequest) error
	GetAllSupplyRequests() ([]models.SupplyRequest, error)
	GetSupplyRequestById(id uint) (*models.SupplyRequest, error)
	GetSupplyRequestsByBusinessId(businessID uint) ([]models.SupplyRequest, error)
	UpdateSupplyRequest(req *models.SupplyRequest) error
	DeleteSupplyRequest(id uint) error
}