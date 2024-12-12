package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecretKey = []byte{}

func InitJWTSercretKey(jwt []byte) {
	jwtSecretKey = jwt
	fmt.Printf("Secret key initialized: %x\n", jwtSecretKey) // для отладки
}

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userID,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// fmt.Printf("Validating token with secret key: %x\n", jwtSecretKey) // для отладки

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})
}
