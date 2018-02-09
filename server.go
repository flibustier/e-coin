package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
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

	log.Print("Server listening on http://localhost:8008/")

	log.Fatal(http.ListenAndServe(":8008", router))
}
