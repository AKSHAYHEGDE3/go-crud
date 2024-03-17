package routes

import (
	"github.com/AKSHAYHEGDE3/go-crud/controllers"
	"github.com/AKSHAYHEGDE3/go-crud/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/registerUser", controllers.RegisterUser)
		auth.POST("/login", controllers.LoginUser)
		auth.GET("/validateUser", middleware.ValidateUser, controllers.Validate)
	}
}
