package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/flibustier/e-coin/repository"
)

// errorResponse writes a response with Status Internal Server Error (500) and log error
func errorResponse(w http.ResponseWriter, err error) {
	log.Println(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode("An error occurred.")
}

// GetUsers serves all registered addresses of our wallet
func GetUsers(w http.ResponseWriter, r *http.Request) {
	addresses, err := repository.GetAddresses()
	if err != nil {
		errorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addresses)
}

// GetUserBalance serves the balance of the current user
func GetUserBalance(w http.ResponseWriter, r *http.Request) {
	address, err := repository.GetUserAddressFromRequest(r)
	if err != nil {
		errorResponse(w, err)
		return
	}

	balances, err := repository.GetBalances(address)
	if err != nil {
		errorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balances)
}

func GetUserAddress(w http.ResponseWriter, r *http.Request) {
	address, err := repository.GetUserAddressFromRequest(r)
	if err != nil {
		errorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(address)
}

// Transaction represents the structure received during a new transaction request
type Transaction struct {
	To     string `json:"to"`
	Assets []struct {
		Name string `json:"name"`
		Qty  int    `json:"qty"`
	} `json:"assets"`
}

// CreateUserTransaction creates a new transaction for the current user
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

	address, err := repository.GetUserAddressFromRequest(r)
	if err != nil {
		errorResponse(w, err)
		return
	}

	for _, asset := range data.Assets {
		if asset.Qty > 0 {
			err = repository.SendAsset(address, data.To, asset.Name, float64(asset.Qty))
			if err != nil {
				errorResponse(w, err)
				break
			}
		}
	}
}

// GetUserTransactions serves the history of all transactions (in and out) for the current user
func GetUserTransactions(w http.ResponseWriter, r *http.Request) {
	// There is no function for it in github.com/golangdaddy/multichain-client, we should implement it!
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]string{})
}
