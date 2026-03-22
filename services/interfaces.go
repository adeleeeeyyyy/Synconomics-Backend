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