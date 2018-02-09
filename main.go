package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golangdaddy/multichain-client"
	"github.com/joho/godotenv"
	"strconv"
)


func main() {
	// Load .env file for setting auth0 secrets and Multichain parameters
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Multichain parameters
	port, err := strconv.Atoi(os.Getenv("MULTICHAIN_PORT"))
	if err != nil {
		log.Fatal("Cannot convert MULTICHAIN_PORT from .env file to integer")
	}

	// Initialize Multichain client
	client = multichain.NewClient(
		os.Getenv("MULTICHAIN_CHAIN_NAME"),
		os.Getenv("MULTICHAIN_RPC_USER"),
		os.Getenv("MULTICHAIN_RPC_PASSWORD"),
		port,
	).ViaNode(
		os.Getenv("MULTICHAIN_HOST"),
		port,
	)

	// Simple command to test everything is fine
	obj, err := client.GetInfo()
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)

	// We can launch the server (finally)!
	StartServer()
}

// The Multichain client previously initialized
var client *multichain.Client









