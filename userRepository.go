package main

import (
	"log"
	"fmt"

	"github.com/boltdb/bolt"
)

const (
	BucketName = "address"		// The name of the bucket for storing addresses
	FileName = "e-coin.db"		// The file name for the data file
)

var db *bolt.DB					// The Bolt Database
var bucket = []byte(BucketName)


func InitializeDatabase() {
	var err error
	// Open the e-coin.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err = bolt.Open(FileName, 0600, nil)
	if err != nil {
		log.Fatal("[FATAL] Could not open data file ", FileName, err)
	}
	log.Printf("[OK] Database successfuly opened in %s data file\n", FileName)

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucket)
		return err
	})
	if err != nil {
		log.Fatal("[FATAL] Could not create bucket ", BucketName, err)
	}
	log.Printf("[OK] Bucket [%s] is ready\n", BucketName)
}


/**
 * This function return the wallet address of the user
 * based on his email address
 * If the user email address doesn't exist in the database,
 * it would be added and a new wallet address will be created
 */
func getUserAddress(email string) (string, error) {
	var address string
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		v := b.Get([]byte(email))
		address = string(v)
		return nil
	})
	if err != nil {
		log.Printf("[ERROR] Could not get address of %s from database", email)
		return "", err
	}
	if address != "" {
		return address, nil
	}
	return address, fmt.Errorf("email doesn't exist")
}

func isUserExisting(email string) bool {
	_, err := getUserAddress(email)
	return err == nil
}


/**
 * This will create a new wallet address and save it in the database at "email" key
 */
func createUser(email string) error {
	// We generate a new wallet address for the user
	address, err := NewAddress()
	if err != nil {
		return err
	}

	// We create an entry in the database with email as key and address as value
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(BucketName))
		if err != nil {
			return err
		}
		return b.Put([]byte(email), []byte(address))
	})
	if err != nil {
		return err
	}

	// We adjust permissions
	Grant(address)

	// We credit the user with the current reward
	CreditAddress(address)

	return nil
}
