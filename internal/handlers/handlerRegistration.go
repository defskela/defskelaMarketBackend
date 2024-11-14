package handlers

import (
	"defskelaMarketBackend/internal/models"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (handler *Handler) Registration(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("UP", user.Password)
	user.Password = string(hashedPassword)
	if err := handler.DB.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			context.JSON(http.StatusOK, gin.H{"message": "Username already exists!"})
			return
		}
		context.JSON(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusOK, user)
	fmt.Println("User registered")
}

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
