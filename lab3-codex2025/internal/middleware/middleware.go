package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method: %s | Path: %s | Time: %s",
			r.Method,
			r.URL.Path,
			time.Now().Format(time.RFC3339),
		)
		next.ServeHTTP(w, r)
	})
}
