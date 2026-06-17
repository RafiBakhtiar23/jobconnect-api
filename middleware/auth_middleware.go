package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"jobconnect-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token diperlukan",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		token, err := utils.ValidateToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token tidak valid",
			})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		fmt.Println("========== JWT CLAIMS ==========")
		fmt.Println(claims)
		fmt.Println("================================")

		role := claims["role"].(string)

		role = strings.TrimSpace(role)

		c.Set("role", role)
		fmt.Println("ROLE DARI JWT =", role)

		c.Next()
	}
}
