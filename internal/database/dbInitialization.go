package database

import (
	"defskelaMarketBackend/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(connectionData string) *gorm.DB {
	var err error
	var db *gorm.DB

	db, err = gorm.Open(postgres.Open(connectionData), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// db.Migrator().DropTable(&models.Cart{})

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Category{},
		&models.Order{}, &models.Cart{}, &models.Market{},
		&models.CartProduct{})
	log.Println("Database connected")
	return db
}
