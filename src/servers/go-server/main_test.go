package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func TestServer(t *testing.T) {
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

	server := httptest.NewServer(corsHandler)
	defer server.Close()

	client := server.Client()

	// Delete All Users
	req, _ := http.NewRequest("DELETE", server.URL+"/users", nil)
	resp, _ := client.Do(req)
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("expected status 204 No Content, got %v", resp.StatusCode)
	}

	// Get All Users (should be empty)
	resp, _ = client.Get(server.URL + "/users")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %v", resp.StatusCode)
	}
	var users []User
	json.NewDecoder(resp.Body).Decode(&users)
	if len(users) != 0 {
		t.Fatalf("expected 0 users, got %v", len(users))
	}

	// Add a User
	user := User{Name: "Test User"}
	body, _ := json.Marshal(user)
	resp, _ = client.Post(server.URL+"/users", "application/json", bytes.NewBuffer(body))
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %v", resp.StatusCode)
	}
	json.NewDecoder(resp.Body).Decode(&user)
	if user.ID != 1 || user.Name != "Test User" {
		t.Fatalf("unexpected user: %v", user)
	}

	// Get User by ID
	resp, _ = client.Get(server.URL + "/users/1")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %v", resp.StatusCode)
	}
	var fetchedUser User
	json.NewDecoder(resp.Body).Decode(&fetchedUser)
	if fetchedUser.ID != 1 || fetchedUser.Name != "Test User" {
		t.Fatalf("unexpected user: %v", fetchedUser)
	}

	// Add Another User
	anotherUser := User{Name: "Another User"}
	body, _ = json.Marshal(anotherUser)
	resp, _ = client.Post(server.URL+"/users", "application/json", bytes.NewBuffer(body))
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %v", resp.StatusCode)
	}
	json.NewDecoder(resp.Body).Decode(&anotherUser)
	if anotherUser.ID != 2 || anotherUser.Name != "Another User" {
		t.Fatalf("unexpected user: %v", anotherUser)
	}

	// Update User by ID
	updatedUser := User{Name: "Updated User"}
	body, _ = json.Marshal(updatedUser)
	req, _ = http.NewRequest("PUT", server.URL+"/users/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %v", resp.StatusCode)
	}
	json.NewDecoder(resp.Body).Decode(&updatedUser)
	if updatedUser.ID != 1 || updatedUser.Name != "Updated User" || updatedUser.HoursWorked != 0 {
		t.Fatalf("unexpected user: %v", updatedUser)
	}

	// Add Hours Worked
	hours := struct {
		Hours int `json:"hours"`
	}{Hours: 5}
	body, _ = json.Marshal(hours)
	req, _ = http.NewRequest("PATCH", server.URL+"/users/1/hours", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %v", resp.StatusCode)
	}
	json.NewDecoder(resp.Body).Decode(&updatedUser)
	if updatedUser.ID != 1 || updatedUser.HoursWorked != 5 {
		t.Fatalf("unexpected user: %v", updatedUser)
	}

	// Add Multiple Users
	for i := 0; i < 9; i++ {
		user := User{Name: "User " + strconv.Itoa(i+3)}
		body, _ := json.Marshal(user)
		resp, _ := client.Post(server.URL+"/users", "application/json", bytes.NewBuffer(body))
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected status 200 OK, got %v", resp.StatusCode)
		}
	}

	// Get All Users
	resp, _ = client.Get(server.URL + "/users")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %v", resp.StatusCode)
	}
	json.NewDecoder(resp.Body).Decode(&users)
	if len(users) != 11 {
		t.Fatalf("expected 11 users, got %v", len(users))
	}

	// Delete a User by ID
	req, _ = http.NewRequest("DELETE", server.URL+"/users/11", nil)
	resp, _ = client.Do(req)
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("expected status 204 No Content, got %v", resp.StatusCode)
	}

	// Verify Deletion
	resp, _ = client.Get(server.URL + "/users/11")
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected status 404 Not Found, got %v", resp.StatusCode)
	}
}
