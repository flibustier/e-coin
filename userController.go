package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func errorResponse(w http.ResponseWriter, err error) {
	log.Println(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode("An error occurred.")
}


/**
 * This action serve all addresses of our wallet
 */
func GetUsers(w http.ResponseWriter, r *http.Request) {
	addresses, err := GetAddresses()
	if err != nil {
		errorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addresses)
}

/**
 * This action serve the balance of the current user, authentication required
 */
func GetUserBalance(w http.ResponseWriter, r *http.Request) {
	// We need the user email address from tokens
	email, err := getUserEmail(r)
	if err != nil {
		errorResponse(w, err)
		return
	}

	// The user is not registered in the database
	if !isUserExisting(email) {
		err = createUser(email)
		if err != nil {
			errorResponse(w, err)
			return
		}
	}

	address, err := getUserAddress(email)
	balances, err := GetBalances(address)
	if err != nil {
		errorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balances)
}

/**
 * This action create an new transaction for the current user
 * Authentication required
 */
func CreateUserTransaction(w http.ResponseWriter, r *http.Request) {

}

/**
 * This action serve the history of all transactions (in and out) for the current user
 * Authentication required
 */
func GetUserTransactions(w http.ResponseWriter, r *http.Request) {
	// There is no function for it in github.com/golangdaddy/multichain-client, we should implement it!
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]string{})
}


