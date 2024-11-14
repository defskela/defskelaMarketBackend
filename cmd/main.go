package main

import (
	"fmt"
	"log"
	"os"

	handlers "defskelaMarketBackend/internal/handlers"

	router "defskelaMarketBackend/internal/router"

	database "defskelaMarketBackend/internal/database"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	connectionData := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db := database.InitDB(connectionData)

	router.InitRouter(handlers.CreateMainHandler(db))
}
