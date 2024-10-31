package handlers

import (
	"fmt"
	"net/http"
	"os"

	"defskelaMarketBackend/internal/models"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func generateJWT(userID uint, secretKey []byte) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // токен действует 24 часа

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

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

	var user models.User
	result := handler.DB.Where("username = ?", input.Username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials user"})
		return
	}

	// Проверка пароля (предполагается, что пароли хранятся в хешированном виде)
	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	// 	return
	// }
	if err := user.Password == input.Password; err != true {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	fmt.Println(user)
	// Здесь следует добавить проверку пароля и существования пользователя

	token, err := generateJWT(user.ID, jwtSecretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	user.Token = token
	handler.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"token": token})
}
