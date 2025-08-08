package main

import (
	"fmt"
	"log"
	"neo-wifi-api/handlers"
	"net/http"
	"os"
)

func main() {
	// Load data from files
	if err := handlers.LoadData(); err != nil {
		log.Fatal("Error loading data:", err)
	}
	port := os.Getenv("PORT")

	// Route handler
	http.HandleFunc("/", handlers.WifiHandler)

	if port == "" {
		port = ":8000"
	}
	fmt.Printf("Server starting on port %s\n", "http://localhost"+port)
	log.Fatal(http.ListenAndServe(port, nil))
}
