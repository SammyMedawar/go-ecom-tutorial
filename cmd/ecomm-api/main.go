package main

import (
	"log"

	"github.com/medawarsammy/go-ecom-tutorial/db"
	"github.com/medawarsammy/go-ecom-tutorial/ecomm-api/handler"
	"github.com/medawarsammy/go-ecom-tutorial/ecomm-api/server"
	"github.com/medawarsammy/go-ecom-tutorial/ecomm-api/storer"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("error opening databse: %v", err)
	}
	//run this when the function ends
	defer db.Close()

	log.Println("Successfully connected to database")
	st := storer.NewMySQLStorer(db.GetDB())
	srv := server.NewServer(st)
	hdl := handler.NewHandler(srv)
	handler.RegisterRoutes(hdl)
	handler.Start(":8080")
}
