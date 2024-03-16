package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/AKSHAYHEGDE3/go-crud/initializers"
	"github.com/AKSHAYHEGDE3/go-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateUser(c *gin.Context) {
	tokenStr, err := c.Cookie("Auth")

	if err != nil {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(401)
		}
		var user models.User
		initializers.DB.First(&user, claims["userid"])

		c.Set("user", user)
		c.Next()

	} else {
		c.AbortWithStatus(401)
	}
}
