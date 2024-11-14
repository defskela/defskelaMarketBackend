package handlers

import (
	"defskelaMarketBackend/internal/models"
	"defskelaMarketBackend/utils"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var emailConfig = utils.EmailConfig{}

func InitEmailConfig(host string, port string, email string, password string) {
	emailConfig.Host = host
	emailConfig.Port = port
	emailConfig.Email = email
	emailConfig.Password = password
}

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

	// Генерируем OTP
	otp, err := utils.GenerateOTP()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"message": "Error creating OTP-code"})
		return
	}

	user.OTP = otp
	user.OTPCreatedAt = time.Now()

	if err := emailConfig.SendEmailOTP(user.Email, otp); err != nil {
		// Если не удалось отправить email, удаляем пользователя
		fmt.Println(err)
		handler.DB.Delete(&user)
	}

	context.JSON(http.StatusOK, user)
	fmt.Println("User registered")
}
