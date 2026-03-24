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
	FindByBusinessID(businessID uint) ([]models.Product, error)
}

type BusinessRepository interface {
	Create(business *models.Business) error
	FindAll() ([]models.Business, error)
	FindByID(id uint) (*models.Business, error)
	Update(business *models.Business) error
	Delete(id uint) error
	FindByUserID(userID uint) ([]models.Business, error)
}

type TransactionRepository interface {
	Create(transaction *models.Transaction) error
	FindAll() ([]models.Transaction, error)
	FindByID(id uint) (*models.Transaction, error)
	Update(transaction *models.Transaction) error
	Delete(id uint) error
	FindByBusinessID(businessID uint) ([]models.Transaction, error)
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
	GetLatestSession(userID, businessID uint, sessionType models.AISessionType) (*models.AISession, error)
}

type SupplyRequestRepository interface {
	Create(req *models.SupplyRequest) error
	FindAll() ([]models.SupplyRequest, error)
	FindByID(id uint) (*models.SupplyRequest, error)
	FindByBusinessID(businessID uint) ([]models.SupplyRequest, error)
	Update(req *models.SupplyRequest) error
	Delete(id uint) error
}

type SupplyOfferRepository interface {
	Create(offer *models.SupplyOffer) error
	FindAll() ([]models.SupplyOffer, error)
	FindByID(id uint) (*models.SupplyOffer, error)
	FindByBusinessID(businessID uint) ([]models.SupplyOffer, error)
	Update(offer *models.SupplyOffer) error
	Delete(id uint) error
}

type SupplyMatchRepository interface {
	Create(match *models.SupplyMatch) error
	FindAll() ([]models.SupplyMatch, error)
	FindByID(id uint) (*models.SupplyMatch, error)
	FindByRequestID(requestID uint) ([]models.SupplyMatch, error)
	FindByOfferID(offerID uint) ([]models.SupplyMatch, error)
	Update(match *models.SupplyMatch) error
	Delete(id uint) error
}

type ThreadRepository interface {
	Create(thread *models.Thread) error
	FindAll() ([]models.Thread, error)
	FindByID(id uint) (*models.Thread, error)
	Update(thread *models.Thread) error
	Delete(id uint) error
}

type ReplyRepository interface {
	Create(reply *models.Reply) error
	FindAll() ([]models.Reply, error)
	FindByID(id uint) (*models.Reply, error)
	FindByThreadID(threadID uint) ([]models.Reply, error)
	Update(reply *models.Reply) error
	Delete(id uint) error
}

type ProductSearchLogRepository interface {
	Create(log *models.ProductSearchLog) error
	FindAll() ([]models.ProductSearchLog, error)
	FindByID(id uint) (*models.ProductSearchLog, error)
	FindByUserID(userID uint) ([]models.ProductSearchLog, error)
}

type MarketTrendRepository interface {
	Create(trend *models.MarketTrend) error
	FindAll() ([]models.MarketTrend, error)
	FindByID(id uint) (*models.MarketTrend, error)
	Update(trend *models.MarketTrend) error
	Delete(id uint) error
	FindTopTen() ([]models.MarketTrend, error)
}

type BusinessMetricRepository interface {
	Create(metric *models.BusinessMetric) error
	FindAll() ([]models.BusinessMetric, error)
	FindByID(id uint) (*models.BusinessMetric, error)
	FindByBusinessID(businessID uint) ([]models.BusinessMetric, error)
	Update(metric *models.BusinessMetric) error
	Delete(id uint) error
	GetLatestByBusinessID(businessID uint) (*models.BusinessMetric, error)
}