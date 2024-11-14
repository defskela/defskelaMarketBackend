package handlers

import (
	"defskelaMarketBackend/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Получить всех юзеров
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
