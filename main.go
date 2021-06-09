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

const connStr string = "host=localhost user=postgres password=Weight!123 dbname=weight_loss-db port=5432 sslmode=disable"

func main() {

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Weight_log{})

	http.HandleFunc("/users", ListUsers)
	http.HandleFunc("/userWeight", ListUserWeightLogs)
	http.ListenAndServe(":3000", nil)

}

// ListUsers: Prints out a list of all the users
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

//ListUserWeightLogs: Prints out a list of all weightlogs of specific user
func ListUserWeightLogs(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	weightLogs := []models.Weight_log{}
	err = db.Where("userid = ?", 1).Find(&weightLogs).Error
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(weightLogs)
	fmt.Println(weightLogs)
}
