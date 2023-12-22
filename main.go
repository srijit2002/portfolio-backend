package main

import (
	"log"
	"net/http"
	"portfolio-backend/routers"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // adjust this to suit your needs
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Check if it's a preflight request and handle it
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass down the request to the next handler (or middleware)
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := routers.Router()
	log.Fatal(http.ListenAndServe(":8000", enableCORS(router)))
}
