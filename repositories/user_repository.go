package repositories

import (
	"Synconomics/models"

	"gorm.io/gorm"
)

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db}
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
    var user models.User
    err := r.db.Where("email = ?", email).First(&user).Error
    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (r *userRepository) FindByGoogleID(googleID string) (*models.User, error) {
    var user models.User
    err := r.db.Where("google_id = ?", googleID).First(&user).Error
    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
    var user models.User
    err := r.db.First(&user, id).Error
    return &user, err
}

func (r *userRepository) Create(user *models.User) error {
    return r.db.Create(user).Error
}

func (r *userRepository) Update(user *models.User) error {
    return r.db.Save(user).Error
}