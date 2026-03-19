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