package handlers

import (
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func CreateMainHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}
