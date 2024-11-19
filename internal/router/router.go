package router

import (
	"defskelaMarketBackend/internal/handlers"
	"defskelaMarketBackend/internal/middlewares"

	_ "defskelaMarketBackend/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title My API
// @version 1.0
// @description This is a sample server for My API.
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @host localhost:8080

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token

func InitRouter(handler *handlers.Handler) {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to MyShop API!"})
	})
	router.GET("/users", handler.GetAllUsers)
	router.Use(middlewares.AuthMiddleware())

	router.HandleMethodNotAllowed = true

	auth := router.Group("/auth")
	{
		auth.POST("/registration", handler.Registration)
		auth.POST("/otp-code", handler.IsTrueOTP)
		auth.POST("/login", handler.Login)
	}

	// router.GET("/users", handler.GetAllUsers)
	router.GET("/markets", handler.GetAllMarkets)
	router.POST("/createMarket", handler.CreateMarket)
	router.GET("/products", handler.GetAllProducts)
	router.POST("/createProduct", handler.CreateProduct)
	router.GET("/products/:market_id", handler.GetProductsByMarketID)

	router.Run(":8080")
}
