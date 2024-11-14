package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string         `json:"username" gorm:"uniqueIndex"`
	Password string         `json:"password"`
	Email    string         `json:"email"`
	Token    pq.StringArray `json:"token" gorm:"type:text[]"`
}
