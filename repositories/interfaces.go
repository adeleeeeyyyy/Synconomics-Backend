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