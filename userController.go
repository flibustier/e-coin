package main

import (
	"encoding/json"
	"io/ioutil"
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
 * This action serve the balance of the current user
 * Authentication required
 */
func GetUserBalance(w http.ResponseWriter, r *http.Request) {
	address, err := getUserAddressFromRequest(r)
	if err != nil {
		errorResponse(w, err)
		return
	}

	balances, err := GetBalances(address)
	if err != nil {
		errorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balances)
}

type Transaction struct{
	To     string  `json:"to"`
	Assets []struct {
		Name string `json:"name"`
		Qty  int    `json:"qty"`
	} `json:"assets"`
}

/**
 * This action create an new transaction for the current user
 * Authentication required
 */
func CreateUserTransaction(w http.ResponseWriter, r *http.Request) {
	var data Transaction
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ERROR] Could not read request body")
		errorResponse(w, err)
		return
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("[ERROR] JSON is invalid in the request body for a transaction")
		errorResponse(w, err)
		return
	}

	address, err := getUserAddressFromRequest(r)
	if err != nil {
		errorResponse(w, err)
		return
	}

	for _, asset := range data.Assets {
		if asset.Qty > 0 {
			err = SendAsset(address, data.To, asset.Name, float64(asset.Qty))
			if err != nil {
				errorResponse(w, err)
				break
			}
		}
	}
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


