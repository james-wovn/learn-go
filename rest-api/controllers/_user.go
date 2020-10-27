package controllers

import (
	"github.com/gin-gonic/gin"
	"go-learn/config"
	"go-learn/models"
	"net/http"
)

type CreateUserInput struct {
	//Name string `json:"name" binding:"required"`
	Email string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Properties
}

type UpdateUserInput struct {
	Name string `json:"name"`
}

// GET /user
// Find all users
func FindUsers(c *gin.Context) {
	var users []models.User

	config.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// POST /user
// Create new user
func CreateUser(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := models.User{
		Name: input.Name
		DisplayLanguage: input.Di
	}
	config.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /user/:id
// Find a user
func FindUser(c *gin.Context) {
	// Get user if exist
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PATCH /user/:id
// Update a user
func UpdateUser(c *gin.Context) {
	// Get user if exist
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// Validate input
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update user
	config.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /user/:id
// Delete a user
func DeleteUser(c *gin.Context) {
	// Get user if exist
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// Delete user
	config.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
