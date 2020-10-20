package main

import (
	"github.com/gin-gonic/gin"
	"go-learn/controllers"
	"go-learn/models"
)

func main() {
	r := gin.Default()

	// Open database connection
	models.OpenDatabase()

	// Routes
	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.FindUser)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	// Run the server
	r.Run()
}
