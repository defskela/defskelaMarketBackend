package handlers

import (
	"defskelaMarketBackend/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Все категории
// @Description Данный запрос позволяет получить список всех категорий
// @Tags categories
// @Accept json
// @Produce json
// @Router /categories [get]
func (handler *Handler) GetAllCategories(context *gin.Context) {
	var categories []models.Category
	handler.DB.Find(&categories)
	if len(categories) == 0 {
		context.JSON(http.StatusOK, gin.H{"message": "Категории не найдены"})
		return
	}
	context.JSON(http.StatusOK, categories)
}

type categoriesArray struct {
	Categories []models.Category `json:"categories" binding:"required"`
}

// @Summary Создание категории
// @Description Данный запрос позволяет создать категорию, если формат данных соответствует структуре models.Category
// @Tags categories
// @Accept json
// @Produce json
// @Param categories body categoriesArray true "Categories data"
// @Router /createCategories [post]
func (handler *Handler) CreateCategories(context *gin.Context) {
	var categoriesData categoriesArray
	if err := context.ShouldBindJSON(&categoriesData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Создаем все категории в транзакции
	tx := handler.DB.Begin()
	for _, category := range categoriesData.Categories {
		if err := tx.Create(&category).Error; err != nil {
			tx.Rollback()
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	tx.Commit()

	context.JSON(http.StatusOK, gin.H{
		"message":    fmt.Sprintf("Successfully created %d categories", len(categoriesData.Categories)),
		"categories": categoriesData.Categories,
	})
}
