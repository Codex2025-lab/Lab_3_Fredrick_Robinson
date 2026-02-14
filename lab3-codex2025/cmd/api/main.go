package main

import (
	"log"
	"net/http"

	"lab3-codex2025/internal/middleware"
	"lab3-codex2025/internal/routes"
)

func main() {
	mux := routes.SetupRoutes()

	log.Println("Server running on port 8080")

	err := http.ListenAndServe(":8080", middleware.LoggingMiddleware(mux))
	if err != nil {
		log.Fatal(err)
	}
}
