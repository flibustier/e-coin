package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	port = 8008
)

func StartServer() {
	// Set routes
	router := mux.NewRouter()

	// For User Controller
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.Handle("/user/balance", authMiddleware(GetUserBalance)).Methods("GET")
	router.HandleFunc("/user/transfer", CreateUserTransaction).Methods("POST")
	router.HandleFunc("/user/transactions", GetUserTransactions).Methods("GET")

	// And for frontend
	router.Handle("/", http.FileServer(http.Dir("./frontend/")))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	log.Printf("[OK] Server listening on http://localhost:%d/", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
