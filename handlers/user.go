package handlers

import (
	"fmt"
	"net/http"
)

// CreateUser - Using the user model create a new user in the DB
func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New User Endpoint Hit")
}

// DeleteUser - Delete a specific user from the DB
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User Endpoint Hit")
}

// UpdateUser - Updates the specific users`s information 
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User Endpoint Hit")
}

// ListAllUsers - Lists out all the users in the DB
func ListAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List All Users Endpoint Hit")
}

