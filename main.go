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

//Prints out a list of all the users
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

//To do Must have a function to ask for daily weight log
func AskWeightLog(which_user int64, weight_log int64) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.Create(&models.Weight_log{
		Weight_log: weight_log,
		Userid:     which_user})
}

//Lists all the weight logs of the specified user
//To do The userID mus not be hardcoded!
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

// The function finds the average in an array
// May be use it inside AverageWeight?
func SumAverage(array []int) int {
	sum := 0
	for _, v := range array {
		sum += v
	}
	result := sum / len(array)
	return result
}

//To do Must have a function for calculating the average weight loss ot gain

func AverageWeight(whichUserID int64) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var allWeightLogs []int
	err = db.Where("userid = ?", whichUserID).Find(&allWeightLogs).Error
	if err != nil {
		log.Println(err)
		return
	}
}

//To do Must have a function for cheering up on consecutive daily logs

//To do Must have a function for cheering up on weight loss

//To do Must have a function for warning if loosing weight too fast

//To do Must have a function for warning if gaining weight instead of loosing

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
