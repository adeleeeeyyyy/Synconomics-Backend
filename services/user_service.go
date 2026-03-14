package services

import (
    "errors"

    "Synconomics/models"
    "Synconomics/pkg"
    "Synconomics/repositories"
    "gorm.io/gorm"
)

type UserService struct {
    repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) Register(req models.RegisterRequest) (*models.User, error) {
    // cek email sudah terdaftar
    _, err := s.repo.FindByEmail(req.Email)
    if err == nil {
        return nil, errors.New("email sudah terdaftar")
    }
    if !errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, err
    }

    hashed, err := pkg.HashPassword(req.Password)
    if err != nil {
        return nil, err
    }

    user := &models.User{
        Name:     req.Name,
        Email:    req.Email,
        Password: hashed,
    }

    if err := s.repo.Create(user); err != nil {
        return nil, err
    }
    return user, nil
}

func (s *UserService) Login(req models.LoginRequest) (string, error) {
    user, err := s.repo.FindByEmail(req.Email)
    if err != nil {
        return "", errors.New("email atau password salah")
    }

    if !pkg.CheckPassword(user.Password, req.Password) {
        return "", errors.New("email atau password salah")
    }

    token, err := pkg.GenerateToken(user.ID)
    if err != nil {
        return "", err
    }
    return token, nil
}

func (s *UserService) GetProfile(id uint) (*models.User, error) {
    return s.repo.FindByID(id)
}