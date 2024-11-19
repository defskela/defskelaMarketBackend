package handlers

import (
	"fmt"
	"net/http"
	"os"

	"defskelaMarketBackend/internal/models"

	"defskelaMarketBackend/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func (handler *Handler) Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := new(models.User)
	result := handler.DB.Where("username = ?", input.Username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials user"})
		return
	}

	fmt.Println("UP:", user.Password, "IP:", input.Password)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials (password)"})
		return
	}
	// if err := user.Password == input.Password; err != true {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	// 	return
	// }
	fmt.Println(user)
	// Здесь следует добавить проверку пароля и существования пользователя

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	user.Token = token
	handler.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"token": token})
}
