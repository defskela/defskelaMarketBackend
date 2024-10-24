package models

import "gorm.io/gorm"

type Market struct {
	gorm.Model
	Name string
	City string
}
