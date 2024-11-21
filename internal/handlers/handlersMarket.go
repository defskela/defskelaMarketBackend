package handlers

import (
	"defskelaMarketBackend/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Получить все продукты
func (handler *Handler) GetAllMarkets(context *gin.Context) {
	var markets []models.Market
	handler.DB.Find(&markets)
	if len(markets) == 0 {
		context.JSON(http.StatusOK, gin.H{"message": "Магазины не найдены"})
		return
	}
	context.JSON(http.StatusOK, markets)
}

// Добавить новый магазин
func (handler *Handler) CreateMarket(context *gin.Context) {
	var market models.Market
	if err := context.ShouldBindJSON(&market); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Неверная структура данных (CreateMarket)"})
		return
	}
	handler.DB.Create(&market)
	context.JSON(http.StatusOK, market)
}

// @Summary Продукты по ID
// @Description Получить список продуктов по marketID
// @Tags product
// @Accept json
// @Produce json
// @Param market_id path int true "Market ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /products/{market_id} [get]
func (handler *Handler) GetProductsByMarketID(context *gin.Context) {
	var products []models.Product
	marketID := context.Param("market_id")

	if err := handler.DB.Where("market_id = ?", marketID).Find(&products).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Магазин с таким id не найден"})
		return
	}

	if len(products) == 0 {
		context.JSON(http.StatusOK, gin.H{"message": "В магазине с таким id не найдены продукты" + marketID})
		return
	}
	context.JSON(http.StatusOK, products)
}
