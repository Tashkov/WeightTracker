package main

import (
	"fmt"
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
	fmt.Println("connect to db")
	//User
	r.GET("/users", handlers.ListAllUsers)
	r.POST("/users", handlers.CreateUser)
	r.GET("/users/:id", handlers.FindUser)
	r.PATCH("/users/:id", handlers.UpdateUser)
	r.DELETE("users/:id", handlers.DeleteUser)

	//Weight
	r.POST("/weight", handlers.CreateWeight)
	r.GET("/weights", handlers.ListAllWeight)
	r.GET("/weights/:id", handlers.FetchUserWeights)

	r.Run()

}
