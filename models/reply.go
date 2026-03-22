package models

import "gorm.io/gorm"

type Reply struct {
	gorm.Model
	ThreadID uint   `json:"thread_id"`
	UserID   uint   `json:"user_id"`
	Thread   Thread `json:"-"`
	User     User   `json:"-"`

	Content string `json:"content" gorm:"type:text"`
}
