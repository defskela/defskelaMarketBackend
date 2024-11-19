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

type registrationData struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type OTPRequest struct {
	OTP string `json:"otp" binding:"required"`
}

// @Summary Код подтверждения
// @Description Для ввода кода подтверждения, который приходит на почту
// @Tags auth
// @Accept json
// @Produce json
// @Param OTPRequest body OTPRequest true "OTPRequest data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /auth/otp-code [post]
// @Security BearerAuth
func (handler *Handler) IsTrueOTP(context *gin.Context) {
	var otp OTPRequest
	if err := context.ShouldBindJSON(&otp); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получаем user_id напрямую
	userID, exists := context.Get("user_id")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	user := new(models.User)
	result := handler.DB.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	fmt.Println("User OTP:", user.OTP, "Input OTP:", otp.OTP)
	user.IsActive = true
	if err := handler.DB.Save(&user).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "user activated",
		"user":    user,
	})
	return
	// Добавьте здесь логику проверки OTP
}

// @Summary Регистрация пользователя
// @Description Для регистрации ребуется передать уникальный username, пароль и уникальный email
// @Tags users, auth
// @Accept json
// @Produce json
// @Param registrationData body registrationData true "registrationData data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /auth/registration [post]
// @Security BearerAuth
func (handler *Handler) Registration(context *gin.Context) {
	// 1. Валидация входных данных
	var requestData registrationData
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestData.Password), bcrypt.DefaultCost)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing password"})
		return
	}

	user := models.User{
		Username: requestData.Username,
		Email:    requestData.Email,
		Password: string(hashedPassword),
		IsActive: false,
	}

	if err := handler.DB.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			context.JSON(http.StatusConflict, gin.H{"error": "Username or email already exists"})
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	otp, err := utils.GenerateOTP()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating OTP code"})
		return
	}

	user.Token = token
	user.OTP = otp
	user.OTPCreatedAt = time.Now()

	if err := handler.DB.Save(&user).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	if err := emailConfig.SendEmailOTP(user.Email, otp); err != nil {
		log.Printf("Failed to send OTP email: %v", err)
		context.JSON(http.StatusOK, gin.H{
			"message": "User registered but email verification failed",
			"user_id": user.ID,
			"token":   token,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Registration successful. Please check your email for verification code",
		"user_id": user.ID,
		"token":   token,
	})
}
