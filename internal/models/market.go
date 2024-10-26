package models

import "gorm.io/gorm"

type Market struct {
	gorm.Model
	Name string `json:"name"`
	City string `json:"city"`
}
