package handlers

import (
	"fmt"
	"net/http"

	"example.com/mod/models"
	"github.com/gin-gonic/gin"
)

// Schema to validate the user input
type CreateUserInput struct {
	First_name string `json:"First_name" binding:"required"`
	Last_name  string `json:"Last_name" binding:"required"`
	Sex        string `json:"Sex" binding:"required"`
	Age        int64  `json:"Age"`
	Height_cm  int64  `json:"Height_cm"`
}

// CreateUser - Using the user model create a new user in the DB
func CreateUser(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create User
	user := models.User{
		First_name: input.First_name,
		Last_name:  input.Last_name,
		Sex:        input.Sex,
		Age:        input.Age,
		Height_cm:  input.Height_cm}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DeleteUser - Delete a specific user from the DB
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Endpoint Hit")
}

// UpdateUser - Updates the specific users`s information
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User Endpoint Hit")
}

// ListAllUsers - Get all users
func ListAllUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
