package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	HoursWorked int    `json:"hoursWorked"`
}

var users = []User{}
var nextID = 1

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, user := range users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	user.ID = nextID
	nextID++
	users = append(users, user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func updateUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, user := range users {
		if user.ID == id {
			var updatedUser User
			json.NewDecoder(r.Body).Decode(&updatedUser)
			updatedUser.ID = id // Ensure the ID remains the same
			users[i] = updatedUser
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func addHoursWorked(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var hours struct {
		HoursToAdd int `json:"hoursToAdd"`
	}
	json.NewDecoder(r.Body).Decode(&hours)
	for i, user := range users {
		if user.ID == id {
			users[i].HoursWorked += hours.HoursToAdd
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func deleteUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func deleteAllUsers(w http.ResponseWriter, r *http.Request) {
	users = []User{}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUserByID).Methods("GET")
	r.HandleFunc("/users", addUser).Methods("POST")
	r.HandleFunc("/users/{id}", updateUserByID).Methods("PUT")
	r.HandleFunc("/users/{id}/hours", addHoursWorked).Methods("PATCH")
	r.HandleFunc("/users/{id}", deleteUserByID).Methods("DELETE")
	r.HandleFunc("/users", deleteAllUsers).Methods("DELETE")

	// Enable CORS for all origins and methods
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)

	http.ListenAndServe(":5004", corsHandler)
}
