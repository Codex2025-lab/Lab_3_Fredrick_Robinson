package routes

import (
	"net/http"
	"lab3-codex2025/internal/handlers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/students", handlers.WriteJSONHandler) // POST
	mux.HandleFunc("/students/all", handlers.ReadJSONHandler) // GET

	return mux
}
