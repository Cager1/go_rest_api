package main

import (
	"github.com/gin-gonic/gin"
	"tim_go_api/controllers"
	"tim_go_api/database"
)

func main() {

	// Connection to database
	database.ConnectToDatabase()

	// Router setup
	router := gin.Default()

	router.GET("/books", controllers.GetBooks)
	router.GET("/book/:id", controllers.GetBook)
	router.POST("/books", controllers.PostBook)
	router.POST("/science", controllers.PostSciences)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
