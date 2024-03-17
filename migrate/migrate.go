package main

import (
	"github.com/AKSHAYHEGDE3/go-crud/initializers"
	"github.com/AKSHAYHEGDE3/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Transaction{})
}
