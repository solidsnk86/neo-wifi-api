package main

import (
	"fmt"
	"log"
	"net/http"
	"wifi-api/handlers"
)

func main() {
	// Cargar datos al inicio
	if err := handlers.LoadData(); err != nil {
		log.Fatal("Error loading data:", err)
	}

	// Setup routes
	http.HandleFunc("/", handlers.WifiHandler)

	// Start server
	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
