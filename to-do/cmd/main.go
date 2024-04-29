package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/prorealize/to-do/api"
	"github.com/prorealize/to-do/database"
	"log"
	"os"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Will use default values or environment variables.")
	}
	// Connects to the database
	database.ConnectDB()

	// Set the router
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	address := fmt.Sprintf("%s:%s", host, port)
	router := api.GetRouter()

	// Start the server
	msg := fmt.Sprintf("Server starting on host %s port %s...", host, port)
	log.Println(msg)
	if err := router.Run(address); err != nil {
		log.Fatal("Server error:", err)
	}
}
