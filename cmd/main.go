package main

import (
	_ "defskelaMarketBackend/docs"
	"defskelaMarketBackend/internal/database"
	"defskelaMarketBackend/internal/handlers"
	"defskelaMarketBackend/internal/router"
	"defskelaMarketBackend/utils"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Server starting...")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	handlers.InitEmailConfig(os.Getenv("SMTPHOST"), os.Getenv("SMTPPORT"), os.Getenv("SENDER_EMAIL"), os.Getenv("SENDER_PASS"))
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	fmt.Printf("Loading secret key: %x\n", secretKey) // для отладки
	utils.InitJWTSercretKey(secretKey)
	connectionData := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db := database.InitDB(connectionData)

	router.InitRouter(handlers.CreateMainHandler(db))
}
