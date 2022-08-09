package middleware

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		c.Abort()
		return
	}
	// Get token
	splittedToken := strings.Split(token, " ")[1]

	parsedToken, err := jwt.ParseWithClaims(splittedToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		c.Abort()
		return
	}

	if !parsedToken.Valid {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		c.Abort()
		return
	}

	c.Next()
}
