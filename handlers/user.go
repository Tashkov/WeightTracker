package handlers

import (
	"net/http"

	"example.com/mod/models"
	"github.com/gin-gonic/gin"
)

// ListAllUsers - Get all users
func ListAllUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// Schema to validate the user input
type CreateUserInput struct {
	FirstName     string `json:"first_name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	Sex           string `json:"sex" binding:"required"`
	Age           int64  `json:"age" binding:"required"`
	Height        int64  `json:"height" binding:"required"`
	CurrentWeight int64  `json:"current_weight" binding:"required"`
	WeightGoal    int64  `json:"weight_goal" binding:"required"`
}

// POST /users
// CreateUser - Using the user model create a new user in the DB
func CreateUser(c *gin.Context) {
	// Validate input
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&input)

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// GET /users/:id
// FindUser - Gets a specific user
func FindUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Debug().Preload("WeightLogs").Where("id=?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Schema to validate the user input
// There was an index out of reach error while trying to PATCH
// Was resolved after adding ID uint to the validation schema
type UpdateUserInput struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Sex           string `json:"sex"`
	Age           int64  `json:"age"`
	Height        int64  `json:"height"`
	CurrentWeight int64  `json:"current_weight"`
	WeightGoal    int64  `json:"weight_goal"`
}

// PATCH /users/:id
// UpdateUser - Updates the specific users`s information
func UpdateUser(c *gin.Context) {
	// Get model if exists
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	// Validate input
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /users/:id
// DeleteUser = Deletes a specific user
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
