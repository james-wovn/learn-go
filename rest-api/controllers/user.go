package controllers

import (
	"github.com/gin-gonic/gin"
	"go-learn/models"
	"net/http"
)

type CreateUserInput struct {
	Name string `json:"name" binding:"required"`
}

func FindUsers(c *gin.Context) {
	var users []models.User

	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
