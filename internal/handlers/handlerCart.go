package handlers

import (
	"defskelaMarketBackend/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type addCartRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
}

// @Summary Добавление товара в корзину
// @Description Требуется авторизация для работы с данным обработчиком для получения cart по user_id из токена
// @Tags cart
// @Accept json
// @Produce json
// @Param addCartRequest body addCartRequest true "addCartRequest data"
// @Router /addProductToCart [post]
// @Security BearerAuth
func (handler *Handler) AddProductToCart(context *gin.Context) {
	var cartReq addCartRequest
	if err := context.ShouldBindJSON(&cartReq); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Неверная структура данных (addProductToCart)"})
		return
	}

	userID, exists := context.Get("user_id")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "id не содержится в токене"})
		return
	}

	user := new(models.User)
	result := handler.DB.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	product := new(models.Product)
	fmt.Println("P ID", cartReq.ProductID)
	result = handler.DB.Where("id = ?", cartReq.ProductID).First(&product)
	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Продукт не найден"})
		return
	}

	var cart models.Cart
	result = handler.DB.Preload("Products").Preload("CartProducts").Where("user_id = ?", userID).First(&cart)
	if result.Error != nil {
		cart = models.Cart{
			UserID: userID.(uint),
		}
		if err := handler.DB.Create(&cart).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания корзины"})
			return
		}
	}

	var cartProduct models.CartProduct
	result = handler.DB.Where("cart_id = ? AND product_id = ?", cart.ID, product.ID).First(&cartProduct)

	if result.Error == nil {
		cartProduct.Quantity++
		handler.DB.Save(&cartProduct)
	} else {
		cartProduct = models.CartProduct{
			CartID:    cart.ID,
			ProductID: product.ID,
			Quantity:  1,
		}
		handler.DB.Create(&cartProduct)
		handler.DB.Model(&cart).Association("Products").Append(product)
	}

	handler.DB.Preload("Products").Preload("CartProducts").First(&cart, cart.ID)

	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Продукт с id %d успешно добавлен в корзину", cartReq.ProductID),
		"cart":    cart.CartProducts,
	})
}
