package handlers

import (
	"defskelaMarketBackend/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) Registration(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	handler.DB.Create(&user)
	context.JSON(http.StatusOK, user)
	fmt.Println("User registered")
}

// Получить все продукты
func (handler *Handler) GetAllUsers(context *gin.Context) {
	var users []models.User
	handler.DB.Find(&users)
	if len(users) == 0 {
		context.JSON(http.StatusOK, gin.H{"message": "No markets found"})
		return
	}
	context.JSON(http.StatusOK, users)
	fmt.Println("Markets fetched")
}
