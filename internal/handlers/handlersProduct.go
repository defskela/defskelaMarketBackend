package handlers

import (
	"defskelaMarketBackend/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Все продукты
// @Description Данный запрос позволяет получить список всех продуктов и данных о них
// @Tags products
// @Accept json
// @Produce json
// @Router /products [get]
func (handler *Handler) GetAllProducts(context *gin.Context) {
	var products []models.Product
	handler.DB.Find(&products)
	if len(products) == 0 {
		context.JSON(http.StatusOK, gin.H{"message": "Продукты не найдены"})
		return
	}
	context.JSON(http.StatusOK, products)
}

type productsArray struct {
	Products []models.Product `json:"products" binding:"required"`
}

// Добавьте новый handler
// @Summary Создание нескольких продуктов
// @Description Создает несколько продуктов за один запрос
// @Tags products
// @Accept json
// @Produce json
// @Param products body productsArray true "Array of products"
// @Router /createProducts [post]
func (handler *Handler) CreateProducts(context *gin.Context) {
	var productsData productsArray
	if err := context.ShouldBindJSON(&productsData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Создаем все продукты в транзакции
	tx := handler.DB.Begin()
	for _, product := range productsData.Products {
		if err := tx.Create(&product).Error; err != nil {
			tx.Rollback()
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	tx.Commit()

	context.JSON(http.StatusOK, gin.H{
		"message":  fmt.Sprintf("Successfully created %d products", len(productsData.Products)),
		"products": productsData.Products,
	})
}
