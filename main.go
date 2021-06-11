package main

import (
	"net/http"

	"example.com/mod/handlers"
	"example.com/mod/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "home page"})
	})

	models.ConnectDataBase() // new

	r.GET("/users", handlers.ListAllUsers)
	r.POST("/new_users", handlers.CreateUser)

	r.Run()

}
