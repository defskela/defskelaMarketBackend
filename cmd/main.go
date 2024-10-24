package main

import (
	"fmt"
	"log"
	"os"

	"defskelaMarketBackend/internal/models"

	handlers "defskelaMarketBackend/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB(connectionData string) {
	var err error
	db, err = gorm.Open(postgres.Open(connectionData), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	log.Println("Database connected")
	db.AutoMigrate(&models.Product{}, &models.Market{})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}
	connectionData := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	initDB(connectionData)

	r := gin.Default()

	handler := handlers.NewHandler(db)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to MyShop API!"})
	})
	r.GET("/markets", handler.GetAllMarkets)
	r.POST("/createMarket", handler.CreateMarket)
	r.GET("/products", handler.GetAllProducts)
	r.POST("/createProduct", handler.CreateProduct)

	r.Run(":8080")
}
