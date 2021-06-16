package handlers

import (
	"net/http"

	"example.com/mod/models"
	"github.com/gin-gonic/gin"
)

// ListAllWeight - Get all weight records
func ListAllWeight(c *gin.Context) {
	var weights []models.Weight_log
	models.DB.Find(&weights)

	c.JSON(http.StatusOK, gin.H{"data": weights})

}

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
		Weight_log: input.Weight_log}
	models.DB.Create(&weight)

	c.JSON(http.StatusOK, gin.H{"data": weight})
}

// GET /weights/:id
// FetchUserWeights - Lists all weight logs of the user
func FetchUserWeights(c *gin.Context) {
	var userWeightLogs []models.Weight_log

	if err := models.DB.Where("User_id=?", c.Param("id")).Find(&userWeightLogs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return

	}

	c.JSON(http.StatusOK, gin.H{"data": userWeightLogs})
}
