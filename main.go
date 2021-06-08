package main

import (
	"Projects/WeightTracker/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// Must have a DB to store the logs

// The function will register a new user
func NewUser(first_name string, last_name string, sex string, age int64, height int64) {
	var err error

	sqlStatement := `
	INSERT INTO users (first_name, last_name, sex, age, height) 
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`
	id := 0
	err = models.DB.QueryRow(sqlStatement, first_name, last_name, sex, age, height).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}

// Must have a function to ask for daily weight log

// Must have a function for calculating the average weight loss ot gain

// Must have a function for cheering up on consecutive daily logs

// Must have a function for cheering up on weight loss

// Must have a function for warning if loosing weight too fast

// Must have a function for warning if gaining weight instead of loosing

func main() {
	var err error

	connStr := "user=postgres dbname=weight_loss-db password=myPass host=localhost sslmode=disable"

	models.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/users", usersIndex)
	http.ListenAndServe(":3000", nil)

	// Commiting to the DB doesn`t seem to work
	// Find out why
	NewUser("Testi", "Testov", "male", 30, 170)
	models.AllUsers()

}

func usersIndex(w http.ResponseWriter, r *http.Request) {
	usrs, err := models.AllUsers()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, usr := range usrs {
		fmt.Fprintf(w, "%s, %s, %s, %d, %d", usr.FirstName, usr.LastName, usr.Sex, usr.Age, usr.HeightCm)
	}

}
