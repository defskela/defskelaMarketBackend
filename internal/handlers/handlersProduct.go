package handlers

import (
	"defskelaMarketBackend/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Получить все продукты
func (handler *Handler) GetAllProducts(context *gin.Context) {
	var products []models.Product
	handler.DB.Find(&products)
	if len(products) == 0 {
		context.JSON(http.StatusOK, gin.H{"message": "No products found"})
		return
	}
	context.JSON(http.StatusOK, products)
	fmt.Println("Products fetched")
}

// Добавить новый продукт
func (handler *Handler) CreateProduct(context *gin.Context) {
	var product models.Product
	if err := context.ShouldBindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	handler.DB.Create(&product)
	context.JSON(http.StatusOK, product)
	fmt.Println("Product created")
}
