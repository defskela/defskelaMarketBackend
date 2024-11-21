package handlers

import (
	"defskelaMarketBackend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Получить все продукты
func (handler *Handler) GetAllProducts(context *gin.Context) {
	var products []models.Product
	handler.DB.Find(&products)
	if len(products) == 0 {
		context.JSON(http.StatusOK, gin.H{"message": "Продукты не найдены"})
		return
	}
	context.JSON(http.StatusOK, products)
}

// Добавить новый продукт
func (handler *Handler) CreateProduct(context *gin.Context) {
	var product models.Product
	if err := context.ShouldBindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Неверная структура данных (CreateProduct)"})
		return
	}
	handler.DB.Create(&product)
	context.JSON(http.StatusOK, product)
}
