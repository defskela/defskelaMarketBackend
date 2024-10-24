package handlers

import (
	"defskelaMarketBackend/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Получить все продукты
func (handler *Handler) GetAllMarkets(context *gin.Context) {
	var markets []models.Market
	handler.DB.Find(&markets)
	context.JSON(http.StatusOK, markets)
	fmt.Println("Markets fetched")
}

// Добавить новый магазин
func (handler *Handler) CreateMarket(context *gin.Context) {
	var market models.Market
	if err := context.ShouldBindJSON(&market); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	handler.DB.Create(&market)
	context.JSON(http.StatusOK, market)
	fmt.Println("Product created")
}

// Получить список продуктов по marketID
func (handler *Handler) GetProductsByMarketID(context *gin.Context) {
	var products []models.Product
	marketID := context.Param("id") // Получаем ID магазина из параметров URL

	// Выполняем запрос для получения всех продуктов с данным MarketID
	if err := handler.DB.Where("market_id = ?", marketID).Find(&products).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем список продуктов
	context.JSON(http.StatusOK, products)
}
