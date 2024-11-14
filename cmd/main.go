package main

import (
	"defskelaMarketBackend/internal/database"
	"defskelaMarketBackend/internal/handlers"
	"defskelaMarketBackend/internal/router"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	handlers.InitEmailConfig(os.Getenv("SMTPHOST"), os.Getenv("SMTPPORT"), os.Getenv("SENDER_EMAIL"), os.Getenv("SENDER_PASS"))

	// Тестируем отправку
	// err = emailConfig.SendEmailOTP("defskela@gmail.com", "123456")
	// if err != nil {
	// 	log.Printf("Ошибка отправки: %v", err)
	// 	return
	// }

	log.Println("Письмо успешно отправлено")

	connectionData := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db := database.InitDB(connectionData)

	router.InitRouter(handlers.CreateMainHandler(db))
}
