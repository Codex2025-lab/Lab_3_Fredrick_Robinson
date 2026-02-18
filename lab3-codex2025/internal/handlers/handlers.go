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

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var student Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	Students = append(Students, student)

	data := envelope{
		"student": student,
	}

	err = writeJSON(w, http.StatusCreated, data, nil)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
}


// ReadJSONHandler handles GET requests
func ReadJSONHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	data := envelope{
		"students": Students,
	}

	err := writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
}

