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

// @Summary Создание продукта
// @Description Данный запрос позволяет создать продукт, если формат данных соответствует структуре models.Product
// @Tags products
// @Accept json
// @Produce json
// @Param market body models.Product true "Product data"
// @Router /createProduct [post]
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
