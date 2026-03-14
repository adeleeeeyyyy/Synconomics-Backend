package models

import "time"

type User struct {
    ID        uint      `json:"id"         gorm:"primaryKey;autoIncrement"`
    Name      string    `json:"name"       gorm:"type:varchar(100);not null"`
    Email     string    `json:"email"      gorm:"type:varchar(100);uniqueIndex;not null"`
    Password  string    `json:"-"          gorm:"type:varchar(255);not null"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type RegisterRequest struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}