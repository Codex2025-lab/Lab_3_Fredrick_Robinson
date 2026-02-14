package handlers

import (
	"encoding/json"
	"net/http"
)

// Student struct represents JSON data
type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// In-memory storage (slice)
var Students []Student

// WriteJSONHandler handles POST requests
func WriteJSONHandler(w http.ResponseWriter, r *http.Request) {
	var student Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	Students = append(Students, student)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

// ReadJSONHandler handles GET requests
func ReadJSONHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Students)
}
