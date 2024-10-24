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

// Добавить новый продукт
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
