package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		role, exists := c.Get("role")

		fmt.Println("ROLE DARI TOKEN :", role)

		if !exists {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Role tidak ditemukan",
			})
			c.Abort()
			return
		}

		roleString, ok := role.(string)

		if !ok {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Role bukan string",
			})
			c.Abort()
			return
		}


		fmt.Printf("ROLE RAW = %#v\n", roleString)
		fmt.Printf("PANJANG ROLE = %d\n", len(roleString))

		if roleString != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Akses hanya untuk admin",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
