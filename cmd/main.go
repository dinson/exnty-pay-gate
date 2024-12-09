package main

import (
	"log"
	"net/http"
	"payment-gateway/client"
	"payment-gateway/config"
	"payment-gateway/internal/api"
)

func init() {
	config.Load()
	client.Init()
}

func main() {
	// Set up the HTTP server and routes
	router := api.SetupRouter()

	// Start the server on port 8080
	log.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
