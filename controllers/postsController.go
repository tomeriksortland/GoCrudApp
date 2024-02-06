package controllers

import (
	"gocrudapp/initializers"
	"gocrudapp/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(context *gin.Context) {

	// Make a struct for the data fields getting sendt into the controller
	var body struct {
		Body  string
		Title string
	}

	context.Bind(&body)

	// Create a new post
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	// Store the new post in the Database
	result := initializers.Database.Create(&post)

	// Check for error when trying to store the new post
	if result.Error != nil {
		context.Status(400)
		return
	}

	// Return statuscode and the post created
	context.JSON(201, gin.H{
		"post": post,
	})
}

func PostsIndex(context *gin.Context) {
	var posts []models.Post
	initializers.Database.Find(&posts)

	context.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(context *gin.Context) {
	id := context.Param("id")
	var post models.Post
	initializers.Database.First(&post, id)

	context.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(context *gin.Context) {
	id := context.Param("id")

	var body struct {
		Body  string
		Title string
	}

	context.Bind(&body)

	var post models.Post
	initializers.Database.First(&post, id)

	initializers.Database.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	context.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(context *gin.Context) {
	id := context.Param("id")

	initializers.Database.Delete(&models.Post{}, id)

	context.Status(200)
}
