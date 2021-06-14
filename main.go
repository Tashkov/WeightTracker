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

	//User
	r.GET("/users", handlers.ListAllUsers)
	r.POST("/users", handlers.CreateUser)
	r.GET("/users/:id", handlers.FindUser)
	r.PATCH("/users/:id", handlers.UpdateUser)
	r.DELETE("users/:id", handlers.DeleteUser)

	//Weight
	r.POST("/weight", handlers.CreateWeight)

	
	r.Run()

}
