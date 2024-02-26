package main

import (
	"github.com/AKSHAYHEGDE3/go-crud/controllers"
	"github.com/AKSHAYHEGDE3/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/registerUser", controllers.RegisterUser)
	r.Run()
}
