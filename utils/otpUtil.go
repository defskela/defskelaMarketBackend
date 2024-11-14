package utils

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"

	"github.com/joho/godotenv"
)

type EmailConfig struct {
	Host     string
	Port     string
	Email    string
	Password string
}

func (c *EmailConfig) SendEmailOTP(toEmail, otp string) error {
	fmt.Println("---", c.Host, c.Port, c.Email, c.Password)
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}
	// Формируем адрес сервера
	smtpServer := fmt.Sprintf("%s:%s", c.Host, c.Port)

	// Формируем заголовки письма
	headers := make(map[string]string)
	headers["From"] = c.Email
	headers["To"] = toEmail
	headers["Subject"] = "Код подтверждения регистрации"
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"utf-8\""

	// Формируем тело письма
	body := fmt.Sprintf(`
        <!DOCTYPE html>
        <html>
        <head>
            <meta charset="UTF-8">
            <style>
                .container {
                    font-family: Arial, sans-serif;
                    max-width: 600px;
                    margin: 0 auto;
                    padding: 20px;
                }
                .code {
                    font-size: 32px;
                    font-weight: bold;
                    color: #2b2b2b;
                    text-align: center;
                    padding: 20px;
                    margin: 20px 0;
                    background-color: #f5f5f5;
                    border-radius: 5px;
                }
                .info {
                    color: #666;
                    font-size: 14px;
                    text-align: center;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <h2>Подтверждение регистрации</h2>
                <p>Здравствуйте!</p>
                <p>Для завершения регистрации введите этот код подтверждения:</p>
                <div class="code">%s</div>
                <p class="info">Код действителен в течение 15 минут.<br>
                Если вы не запрашивали этот код, просто проигнорируйте это письмо.</p>
            </div>
        </body>
        </html>`, otp)

	// Собираем сообщение
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Настраиваем TLS конфигурацию
	tlsConfig := &tls.Config{
		ServerName:         c.Host,
		InsecureSkipVerify: false,
	}

	// Устанавливаем соединение
	conn, err := tls.Dial("tcp", smtpServer, tlsConfig)
	if err != nil {
		return fmt.Errorf("failed to create TLS connection: %v (server: %s)", err, smtpServer)
	}
	defer conn.Close()

	// Создаем SMTP клиент
	client, err := smtp.NewClient(conn, c.Host)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %v", err)
	}
	defer client.Close()

	// Аутентификация
	auth := smtp.PlainAuth("", c.Email, c.Password, c.Host)
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("failed to authenticate: %v", err)
	}

	// Отправитель
	if err = client.Mail(c.Email); err != nil {
		return fmt.Errorf("failed to set sender: %v", err)
	}

	// Получатель
	if err = client.Rcpt(toEmail); err != nil {
		return fmt.Errorf("failed to set recipient: %v", err)
	}

	// Записываем тело письма
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to create message writer: %v", err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("failed to write message: %v", err)
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("failed to close writer: %v", err)
	}

	return client.Quit()
}

// GenerateOTP генерирует 6-значный код подтверждения
func GenerateOTP() (string, error) {
	numbers := make([]byte, 6)
	if _, err := rand.Read(numbers); err != nil {
		return "", err
	}

	for i := range numbers {
		numbers[i] = numbers[i]%10 + '0'
	}

	return string(numbers), nil
}
