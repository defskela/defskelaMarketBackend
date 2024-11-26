package handlers

import (
	"defskelaMarketBackend/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Все магазины
// @Description Данный запрос позволяет получить список всех магазинов, их данных и товаров внутри них
// @Tags markets
// @Accept json
// @Produce json
// @Router /markets [get]
func (handler *Handler) GetAllMarkets(context *gin.Context) {
	var markets []models.Market
	handler.DB.Find(&markets)
	if len(markets) == 0 {
		context.JSON(http.StatusOK, gin.H{"message": "Магазины не найдены"})
		return
	}
	context.JSON(http.StatusOK, markets)
}

type marketsArray struct {
	Markets []models.Market `json:"markets" binding:"required"`
}

// @Summary Создание магазина
// @Description Данный запрос позволяет создать магазин, если формат данных соответствует структуре models.Market
// @Tags markets
// @Accept json
// @Produce json
// @Param market body marketsArray true "Markets data"
// @Router /createMarkets [post]
func (handler *Handler) CreateMarkets(context *gin.Context) {
	var marketsData marketsArray
	if err := context.ShouldBindJSON(&marketsData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Создаем все продукты в транзакции
	tx := handler.DB.Begin()
	for _, market := range marketsData.Markets {
		if err := tx.Create(&market).Error; err != nil {
			tx.Rollback()
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	tx.Commit()

	context.JSON(http.StatusOK, gin.H{
		"message":  fmt.Sprintf("Successfully created %d markets", len(marketsData.Markets)),
		"products": marketsData.Markets,
	})
}
