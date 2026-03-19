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
