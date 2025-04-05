package main

import (
	"log"

	"github.com/medawarsammy/go-ecom-tutorial/db"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("error opening databse: %v", err)
	}
	//run this when the function ends
	defer db.Close()

	log.Println("Successfully connected to database")
}
