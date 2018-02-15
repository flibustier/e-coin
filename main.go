package main

import (
	"log"

	"github.com/flibustier/e-coin/repository"
	"github.com/flibustier/e-coin/server"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file for setting auth0 secrets and Multichain parameters
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	repository.InitializeBlockchain()

	repository.InitializeDatabase()

	// We can launch the server (finally)!
	server.StartServer()
}
