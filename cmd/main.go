package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"defskelaMarketBackend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

// Получить все продукты
func (handler *Handler) GetAllProducts(context *gin.Context) {
	var products []models.Product
	handler.DB.Find(&products)
	context.JSON(http.StatusOK, products)
	fmt.Println("Products fetched")
}

// Добавить новый продукт
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

func initDB(connectionData string) {
	var err error
	db, err = gorm.Open(postgres.Open(connectionData), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	log.Println("Database connected")
	db.AutoMigrate(&models.Product{})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}
	connectionData := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	initDB(connectionData)

	r := gin.Default()

	productHandler := NewHandler(db)

	r.GET("/products", productHandler.GetAllProducts)
	r.POST("/createProduct", productHandler.CreateProduct)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to MyShop API!"})
	})

	r.Run(":8080")
}
