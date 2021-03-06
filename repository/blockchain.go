package repository

import (
	"log"
	"math"
	"os"
	"strconv"

	"github.com/flibustier/multichain-client"
)

const (
	InitialReward         = 10     // Quantity of asset that will be given to new users
	RewardName            = "blue" // Asset's name for reward
	RewardHalvingInterval = 100    // Every time the number of addresses in our wallet reach this value, the reward will be halved
	HistoryLength         = 20     // The number of transactions shown for the address history
)

// numberOfAddresses stores the number of address in the wallet to avoid useless API calls
var numberOfAddresses uint32

// client is the Multichain client API initialized
var client *multichain.Client

// currentRewardQuantity calculates the current reward quantity based on the number of addresses and reward halving interval
func currentRewardQuantity() float64 {
	return InitialReward / math.Pow(2, math.Trunc(float64(numberOfAddresses)/RewardHalvingInterval))
}

// CreditAddress credits an address with the current reward quantity
func CreditAddress(address string) {
	client.IssueMore(address, RewardName, currentRewardQuantity())
}

// Grant grants 'send' and 'receive' permissions to the address
func Grant(address string) {
	var permissions []string
	permissions = append(permissions, "receive")
	permissions = append(permissions, "send")
	client.Grant([]string{address}, permissions)
}

// SendAsset sends a quantity of asset name to an address
func SendAsset(from string, to string, name string, value float64) error {
	response, err := client.SendAssetFrom(from, to, name, value)
	if err != nil {
		log.Printf("[ERROR] %f quantity of asset %s couldn't be sent from %s to %s\n", value, name, from, to)
		log.Println(response)
		return err
	}
	return nil
}

// NewAddress generates a new wallet address
func NewAddress() (string, error) {
	response, err := client.GetNewAddress()
	if err != nil {
		log.Println("[ERROR] Could not get a new address from Multichain!")
		return "", err
	}
	numberOfAddresses++
	address := response.Result().(string)
	log.Printf("[INFO] A new address has been generated (%s)\n", address)
	return address, nil
}

// GetBalances returns the asset balance for the address
func GetBalances(address string) (interface{}, error) {
	balance, err := client.GetAddressBalances(address)
	if err != nil {
		log.Printf("[ERROR] Could not get address balance for %s\n", address)
		return nil, err
	}
	return balance.Result(), nil
}

// GetAddresses returns the list of addresses in our wallet
func GetAddresses() (interface{}, error) {
	addresses, err := client.GetAddresses(false)
	if err != nil {
		log.Println("[ERROR] Could not get addresses from Multichain!")
		return nil, err
	}
	return addresses.Result(), nil
}

// GetHistory returns the 'HistoryLength' last transactions for address
func GetHistory(address string) (interface{}, error) {
	history, err := client.ListAddressTransactions(address, HistoryLength, 0, false)
	if err != nil {
		log.Printf("[ERROR] Could not get history from %s address\n", address)
		return nil, err
	}
	return history.Result(), nil
}

func InitializeBlockchain() {
	// Multichain Port parameter conversion
	port, err := strconv.Atoi(os.Getenv("MULTICHAIN_PORT"))
	if err != nil {
		log.Fatal("[FATAL] Cannot convert MULTICHAIN_PORT from .env file to integer")
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

	// Simple command to test everything is fine and giving information of Multichain node
	obj, err := client.GetInfo()
	if err != nil {
		log.Fatal("[FATAL] Cannot get informations from Multichain RPC\n", err)
	}
	log.Println("[OK] Multichain is up and running", obj.Result())

	// Get the number of addresses in our wallet
	obj, err = client.GetAddresses(false)
	if err != nil {
		log.Fatal("[FATAL] Could not get addresses from Multichain", err)
	}
	addresses := obj.Result().([]interface{})
	numberOfAddresses = uint32(len(addresses))
	log.Printf("[OK] Your wallet currently have %d address(es)\n", numberOfAddresses)
	log.Printf("[OK] The current reward is fixed to %f\n", currentRewardQuantity())

	// If it's the first time the node is launched, we have to create the asset for reward
	log.Printf("[OK] Your main address is %s\n", addresses[0])
	// listAssets is not implemented :(
	// that why we try to issue even if it already exists
	obj, err = client.Issue(true, addresses[0].(string), RewardName, InitialReward, 1)
	if err != nil {
		log.Printf("[OK] Asset %s seems already existing", RewardName)
	} else {
		log.Printf("[OK] Asset %s successfuly created", RewardName)
	}

}
