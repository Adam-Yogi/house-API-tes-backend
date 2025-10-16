package main

import (
	"gin-api/package/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.HouseListing(r)

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Starting server on port %s\n", port)

	corsHandler := handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)

	err := http.ListenAndServe(":"+port, corsHandler)
	if err != nil {
		log.Fatal(err)
	}
}
