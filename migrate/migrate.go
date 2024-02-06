package main

import (
	"gocrudapp/initializers"
	"gocrudapp/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.Database.AutoMigrate(&models.Post{})
}
