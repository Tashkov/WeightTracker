package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"example.com/mod/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Must have a DB to store the logs

// The function will register a new user

// Must have a function to ask for daily weight log

// Must have a function for calculating the average weight loss ot gain

// Must have a function for cheering up on consecutive daily logs

// Must have a function for cheering up on weight loss

// Must have a function for warning if loosing weight too fast

// Must have a function for warning if gaining weight instead of loosing

const connStr string = "host=localhost user=postgres password=Weight!123 dbname=weight_loss-db port=5432 sslmode=disable"

func main() {

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Weight{})

	db.Create(&models.User{
		Name: "Hristo",
		Age:  27,
	})
	db.Create(&models.User{
		Name: "Ivancho",
		Age:  20})

	http.HandleFunc("/users", ListUsers)
	http.ListenAndServe(":3000", nil)

}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	usrs := []models.User{}
	err = db.Find(&usrs).Error
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usrs)
	fmt.Println(usrs)

}
