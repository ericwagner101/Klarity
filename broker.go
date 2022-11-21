package main

import (
	"fmt"
	"git.mills.io/prologic/bitcask"
)

func createBitcask() *bitcask.Bitcask {
	// Create a new Bitcask instance
	db, err := bitcask.Open("/tmp/bitcask")
	if err != nil {
		panic(err)
	}
	return db
}

func cleanup() error {
	fmt.Println("Running cleanup...")
	return fmt.Errorf("error on cleanup")
}

func main() {
	db := createBitcask()
	defer db.Close()

	// Store a key
	err := db.Put([]byte("hello"), []byte("world"))
	if err != nil {
		panic(err)
	}

	// Retrieve a key
	value, err := db.Get([]byte("hello"))
	if err != nil {
		panic(err)
	}

	fmt.Println(string(value))
}
