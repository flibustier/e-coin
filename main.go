package main

import (
	"log"

	"github.com/joho/godotenv"
)


func main() {
	// Load .env file for setting auth0 secrets and Multichain parameters
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	InitializeBlockchain()

	InitializeDatabase()
	defer db.Close()

	// We can launch the server (finally)!
	StartServer()
}










