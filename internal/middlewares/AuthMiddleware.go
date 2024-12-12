package middlewares

import (
	"defskelaMarketBackend/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Пропускаем авторизацию для определенных эндпоинтов
		if c.Request.URL.Path == "/auth/login" ||
			c.Request.URL.Path == "/auth/registration" ||
			strings.HasPrefix(c.Request.URL.Path, "/swagger") {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Проверяем формат токена, необходммо, чтобы он содержал Bearer
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		token, err := utils.ValidateToken(bearerToken[1])
		// fmt.Println("TOKEN", token)
		if err != nil {
			fmt.Println("ERR", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token. Auth no validate"})
			c.Abort()
			return
		}

		// Добавляем claims в контекст
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Устанавливаем claims в контекст
			c.Set("claims", claims)
			c.Set("user_id", uint(claims["user_id"].(float64)))
		}

		c.Next()
	}
}
