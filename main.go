package main

import (
	"gocrudapp/controllers"
	"gocrudapp/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	router := gin.Default()

	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostShow)
	router.POST("/posts", controllers.PostsCreate)
	router.PUT("/posts/:id", controllers.PostUpdate)
	router.DELETE("/posts/:id", controllers.DeletePost)
	router.Run()
}
