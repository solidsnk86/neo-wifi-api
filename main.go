package main

import (
	"fmt"
	"log"
	"neo-wifi-api/handlers"
	"net/http"
)

func main() {
	// Cargar datos al inicio
	if err := handlers.LoadData(); err != nil {
		log.Fatal("Error loading data:", err)
	}

	// Routes
	http.HandleFunc("/", handlers.WifiHandler)

	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
