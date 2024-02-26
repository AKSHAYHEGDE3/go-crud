package controllers

import (
	// "github.com/AKSHAYHEGDE3/go-crud/initializers"
	// "github.com/AKSHAYHEGDE3/go-crud/models"
	// import "github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var body struct {
		Email    string
		Password string
		Name     string
	}

	c.Bind(&body)

}
