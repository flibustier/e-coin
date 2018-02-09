package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/**
 * Hardcoded for now!
 * This function will return the wallet address of the user
 */
func getUserAddress() string {
	return "1LHoSVS4q35E9xLMczfuNPznLrraaLM7isLfYG"
}

/**
 * This action serve all addresses of our wallet
 */
func GetUsers(w http.ResponseWriter, r *http.Request) {
	addresses, err := client.GetAddresses(false)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addresses.Result())
}

/**
 * This action serve the balance of the current user, authentication required
 */
var GetUserBalance = http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {

	email, err := getUserEmail(r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(email)
	}

	/*
		Todo: We have the email address, when need now a small database
		to translate (email) to (blockchain wallet address)
	*/

	balance, err := client.GetAddressBalances(getUserAddress())
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balance.Result())
})

/**
 * This action serve the history of all transactions (in and out) for the current user
 * Authentication required
 */
func GetUserTransactions(w http.ResponseWriter, r *http.Request) {
	// There is no function for it in github.com/golangdaddy/multichain-client, we should implement it!
}

/**
 * This action create an new transaction for the current user
 * Authentication required
 */
func CreateUserTransaction(w http.ResponseWriter, r *http.Request) {}


