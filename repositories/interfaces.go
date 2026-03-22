package repositories

import "Synconomics/models"

type UserRepository interface{
	FindByEmail(email string) (*models.User, error)
	FindByGoogleID(googleID string) (*models.User, error)
	FindByID(id uint) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
}

type ProductRepository interface {
	Create(product *models.Product) error
	FindAll() ([]models.Product, error)
	FindByID(id uint) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id uint) error
}

type BusinessRepository interface {
	Create(business *models.Business) error
	FindAll() ([]models.Business, error)
	FindByID(id uint) (*models.Business, error)
	Update(business *models.Business) error
	Delete(id uint) error
}

type TransactionRepository interface {
	Create(transaction *models.Transaction) error
	FindAll() ([]models.Transaction, error)
	FindByID(id uint) (*models.Transaction, error)
	Update(transaction *models.Transaction) error
	Delete(id uint) error
}

type TransactionItemRepository interface {
	Create(item *models.TransactionItem) error
	FindAll() ([]models.TransactionItem, error)
	FindByID(id uint) (*models.TransactionItem, error)
	FindByTransactionID(transactionID uint) ([]models.TransactionItem, error)
	Update(item *models.TransactionItem) error
	Delete(id uint) error
}

type ExpenseRepository interface {
	Create(expense *models.Expense) error
	FindAll() ([]models.Expense, error)
	FindByID(id uint) (*models.Expense, error)
	FindByBusinessID(businessID uint) ([]models.Expense, error)
	Update(expense *models.Expense) error
	Delete(id uint) error
}

type AIRepository interface {
	CreateSession(session *models.AISession) error
	GetSessionByID(id uint) (*models.AISession, error)
	GetSessionsByUserID(userID uint) ([]models.AISession, error)
	SaveMessage(message *models.AIMessage) error
	GetMessagesBySessionID(sessionID uint) ([]models.AIMessage, error)
	SaveResult(result *models.AIResult) error
	GetResultBySessionID(sessionID uint) (*models.AIResult, error)
}

type SupplyRequestRepository interface {
	Create(req *models.SupplyRequest) error
	FindAll() ([]models.SupplyRequest, error)
	FindByID(id uint) (*models.SupplyRequest, error)
	FindByBusinessID(businessID uint) ([]models.SupplyRequest, error)
	Update(req *models.SupplyRequest) error
	Delete(id uint) error
}