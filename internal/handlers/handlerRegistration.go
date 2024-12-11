package handlers

import (
	"defskelaMarketBackend/internal/models"
	"defskelaMarketBackend/utils"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const (
	errEmailInvalid    = "Неверный формат почты"
	errPasswordLength  = "Пароль должен быть длиной не менее 8 символов"
	errPasswordUpper   = "Пароль должен содержать хотя бы одну заглавную букву"
	errPasswordLower   = "Пароль должен содержать хотя бы одну прописную букву"
	errPasswordDigit   = "Пароль должен содержать хотя бы одно число"
	errPasswordSpecial = "Пароль должен содержать хотя бы один специальный символ"
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
// @Router /auth/otp-code [post]
// @Security BearerAuth
func (handler *Handler) IsTrueOTP(context *gin.Context) {
	var otp OTPRequest
	if err := context.ShouldBindJSON(&otp); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Неверная структура данных (IsTrueOTP)"})
		return
	}

	userID, exists := context.Get("user_id")
	fmt.Println(userID)
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Непредвиденная ошибка получения id, повторите попытку"})
		return
	}

	user := new(models.User)
	result := handler.DB.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}
	if otp.OTP == user.OTP {
		user.IsActive = true
		if err := handler.DB.Save(&user).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обработке otp-кода, повторите попытку"})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "user activated",
			"user":    user,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "otp-код не совпадает",
	})
	return
}

// @Summary Регистрация пользователя
// @Description Для регистрации ребуется передать уникальный username, пароль и уникальный email
// @Tags users, auth
// @Accept json
// @Produce json
// @Param registrationData body registrationData true "registrationData data"
// @Router /auth/registration [post]
// @Security BearerAuth
func (handler *Handler) Registration(context *gin.Context) {
	var requestData registrationData
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Неверная структура данных (Registration)"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestData.Password), bcrypt.DefaultCost)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка хеширования пароля"})
		return
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(requestData.Email) {
		context.JSON(http.StatusBadRequest, gin.H{"error": errEmailInvalid})
		return
	}

	password := requestData.Password

	if len(password) < 8 {
		context.JSON(http.StatusBadRequest, gin.H{"error": errPasswordLength})
		return
	}

	upperRegex := regexp.MustCompile(`[A-Z]`)
	if !upperRegex.MatchString(password) {
		context.JSON(http.StatusBadRequest, gin.H{"error": errPasswordUpper})
		return
	}

	lowerRegex := regexp.MustCompile(`[a-z]`)
	if !lowerRegex.MatchString(password) {
		context.JSON(http.StatusBadRequest, gin.H{"error": errPasswordLower})
		return
	}

	digitRegex := regexp.MustCompile(`[0-9]`)
	if !digitRegex.MatchString(password) {
		context.JSON(http.StatusBadRequest, gin.H{"error": errPasswordDigit})
		return
	}

	specialRegex := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)
	if !specialRegex.MatchString(password) {
		context.JSON(http.StatusBadRequest, gin.H{"error": errPasswordSpecial})
		return
	}

	user := models.User{
		Username: requestData.Username,
		Email:    requestData.Email,
		Password: string(hashedPassword),
		IsActive: false,
	}

	if err := handler.DB.Create(&user).Error; err != nil {
		if user.IsActive == true {
			if strings.Contains(err.Error(), "duplicate key value") {
				context.JSON(http.StatusConflict, gin.H{"error": "Username or email already exists"})
				return
			}
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка регистрации"})
		return
	}

	otp, err := utils.GenerateOTP()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка отправки кода на почту"})
		return
	}

	user.Token = token
	user.OTP = otp
	user.OTPCreatedAt = time.Now()

	if err := handler.DB.Save(&user).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сохранения данных"})
		return
	}

	if err := emailConfig.SendEmailOTP(user.Email, otp); err != nil {
		context.JSON(http.StatusConflict, gin.H{
			"message": "User registered but email verification failed",
			"user_id": user.ID,
			"token":   token,
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Registration successful. Please check your email for verification code",
		"otp":     user.OTP,
		"user_id": user.ID,
		"token":   token,
	})
}
