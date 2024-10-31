package router

import (
	"defskelaMarketBackend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func InitRouter(handler *handlers.Handler) {
	router := gin.Default()

	router.HandleMethodNotAllowed = true

	auth := router.Group("/auth")
	{
		auth.POST("/registration", handler.Registration)
		auth.POST("/login", handler.Login)
	}
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to MyShop API!"})
	})
	router.GET("/users", handler.GetAllUsers)
	router.GET("/markets", handler.GetAllMarkets)
	router.POST("/createMarket", handler.CreateMarket)
	router.GET("/products", handler.GetAllProducts)
	router.POST("/createProduct", handler.CreateProduct)
	router.GET("/markets/:id/products", handler.GetProductsByMarketID)

	router.Run(":8080")
}
