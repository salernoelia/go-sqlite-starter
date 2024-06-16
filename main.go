package main

import (
	"go-sqlite/database"
	"go-sqlite/handlers"
	"log"
	"net/http"
)

func main() {
	// Initialize the database
	database.SetupDatabase()

	// Route Handlers
	http.HandleFunc("/api/products", handlers.ProductsHandler)
	http.HandleFunc("/api/products/", handlers.ProductHandler)
	http.HandleFunc("/api/register", handlers.RegisterHandler)
	http.HandleFunc("/api/login", handlers.LoginHandler)

	// Start server
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
