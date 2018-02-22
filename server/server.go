package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flibustier/e-coin/controller"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func homeHandler(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "./frontend/index.html")
}

func StartServer(port uint16) {
	// Set routes
	router := mux.NewRouter()

	// For User Controller
	api := router.PathPrefix("/users").Subrouter().StrictSlash(true)
	api.Methods("GET").Path("").HandlerFunc(controller.GetUsers)
	api.Methods("GET").Path("/balance").HandlerFunc(controller.GetUserBalance)
	api.Methods("GET").Path("/address").HandlerFunc(controller.GetUserAddress)
	api.Methods("POST").Path("/transfer").HandlerFunc(controller.CreateUserTransaction)
	api.Methods("GET").Path("/transactions").HandlerFunc(controller.GetUserTransactions)
	api.Use(authMiddleware)
	api.Use(idMiddleware)

	// And for frontend
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/login", homeHandler)
	router.HandleFunc("/callback", homeHandler)
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./frontend/css/"))))
	router.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir("./frontend/dist/"))))
	router.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("./frontend/fonts/"))))

	host := fmt.Sprintf("http://localhost:%d", port)

	// CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080", "https://flibustier.github.io", host},
		AllowedHeaders: []string{"authorization", "id_token", "content-type"},
	})

	log.Printf("[OK] Server listening on %s\n", host)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), c.Handler(router)))
}
