package handlers

import (
	"defskelaMarketBackend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Все пользователи
// @Description Данный запрос позволяет получить список всех пользователей и их данных
// @Tags users
// @Accept json
// @Produce json
// @Router /users [get]
func (handler *Handler) GetAllUsers(context *gin.Context) {
	var users []models.User
	handler.DB.Find(&users)
	if len(users) == 0 {
		context.JSON(http.StatusOK, gin.H{"message": "Пользователи не найдены"})
		return
	}
	context.JSON(http.StatusOK, users)
}
