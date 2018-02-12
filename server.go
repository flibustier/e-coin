package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	port = 8008
)

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "./frontend/index.html")
}

func StartServer() {
	// Set routes
	router := mux.NewRouter()

	// For User Controller
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/user/balance", GetUserBalance).Methods("GET")
	router.HandleFunc("/user/transfer", CreateUserTransaction).Methods("POST")
	router.HandleFunc("/user/transactions", GetUserTransactions).Methods("GET")

	// And for frontend

	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/login", HomeHandler)
	router.HandleFunc("/callback", HomeHandler)
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./frontend/css/"))))
	router.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir("./frontend/dist/"))))
	router.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("./frontend/fonts/"))))

	// CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080", "https://flibustier.github.io"},
		AllowedHeaders: []string{"authorization", "id_token"},
	})

	log.Printf("[OK] Server listening on http://localhost:%d/", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), c.Handler(authMiddleware(router))))
}
