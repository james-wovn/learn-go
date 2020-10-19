package main

import (
	"github.com/gin-gonic/gin"
	"go-learn/controllers"
	"go-learn/models"
	"net/http"
)

func main() {
	r := gin.Default()

	models.OpenDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/users", controllers.FindUsers)
	r.POST("/user/create", controllers.CreateUser)
	r.GET("/user/:id", controllers.FindUser)

	r.Run()
}
