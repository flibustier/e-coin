package main

import (
	"log"

	"flag"
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

	port := flag.Uint("p", 8008, "port to listen on")
	flag.Parse()

	repository.InitializeBlockchain()

	repository.InitializeDatabase()

	// We can launch the server (finally)!
	server.StartServer(uint16(*port))
}
