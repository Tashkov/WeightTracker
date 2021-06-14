package handlers

import (
	"net/http"

	"example.com/mod/models"
	"github.com/gin-gonic/gin"
)

// Schema to validate the user input
type CreateWeightInput struct {
	Weight_log int64 `json:"Weight_log" binding:"required"`
	User_id    uint  `json:"User_id" binding:"required"`
}

//POST /weight
//CreateWeight - Using the user model create a new user in the DB
func CreateWeight(c *gin.Context) {
	//Validate input
	var input CreateWeightInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Create Weight
	weight := models.Weight_log{
		Weight_log: input.Weight_log,
		User_id:    input.User_id}
	models.DB.Create(&weight)

	c.JSON(http.StatusOK, gin.H{"data": weight})
}
