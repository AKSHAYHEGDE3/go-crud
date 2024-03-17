package main

import (
	"github.com/AKSHAYHEGDE3/go-crud/initializers"
	"github.com/AKSHAYHEGDE3/go-crud/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	routes.AuthRoutes(r)

	r.Run()
}
