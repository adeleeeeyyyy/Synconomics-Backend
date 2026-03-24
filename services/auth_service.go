package services

import (
	"Synconomics/config"
	"Synconomics/models"
	"Synconomics/repositories"
	"errors"

	"github.com/markbates/goth"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthServices(userRepo repositories.UserRepository) AuthService {
	return &authService{userRepo}
}

func (s *authService) Register(name, email, password string) (*models.User, string, error) {
	existing, err := s.userRepo.FindByEmail(email)
	if err == nil && existing != nil {
		return nil, "", errors.New("Email already registered")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, "", err
	}

	pass := string(hashed)
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: &pass,
		Provider: "manual",
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, "", err
	}

	token, err := config.GenerateToken(user.ID)
	return user, token, err
}

func (s *authService) Login(email, password string) (*models.User, string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, "", errors.New("invalid credential")
	}

	if user.Password == nil {
		return nil, "", errors.New("this accound uses google login")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password)); err != nil {
		return nil, "", errors.New("Invalid credentials")
	}

	token, err := config.GenerateToken(user.ID)
	return user, token, err
}

func (s *authService) HandleGoogleCallback(googleUser goth.User) (*models.User, string, error) {
	user, err := s.userRepo.FindByGoogleID(googleUser.UserID)
	if err == nil {
		token, err := config.GenerateToken(user.ID)
		return user, token, err
	}

	existing, err := s.userRepo.FindByEmail(googleUser.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, "", err
	}

	if existing != nil {
		existing.GoogleID = &googleUser.UserID
		existing.Provider = "both"
		if existing.Avatar == nil && googleUser.AvatarURL != "" {
			existing.Avatar = &googleUser.AvatarURL
		}
		s.userRepo.Update(existing)
		token, err := config.GenerateToken(existing.ID)
		return existing, token, err
	}

	avatar := googleUser.AvatarURL
	googleID := googleUser.UserID
	newUser := &models.User{
		Name:     googleUser.Name,
		Email:    googleUser.Email,
		GoogleID: &googleID,
		Avatar:   &avatar,
		Provider: "google",
	}

	if err := s.userRepo.Create(newUser); err != nil {
		return nil, "", err
	}

	token, err := config.GenerateToken(newUser.ID)
	return newUser, token, err
}

func (s *authService) GetProfile(userID uint) (*models.User, string, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, "", err
	}
	token, err := config.GenerateToken(user.ID)
	return user, token, err
}
func (s *authService) UpdateProfile(userID uint, name, email string) (*models.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	if name != "" {
		user.Name = name
	}
	if email != "" {
		// Check if email taken
		existing, err := s.userRepo.FindByEmail(email)
		if err == nil && existing != nil && existing.ID != userID {
			return nil, errors.New("email already taken")
		}
		user.Email = email
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}
