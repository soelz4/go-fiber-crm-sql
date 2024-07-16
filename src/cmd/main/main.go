package main

import (
	"fmt"
	"log"

	"go-fiber-crm-sql/src/cmd/server"
	"go-fiber-crm-sql/src/db"
	"go-fiber-crm-sql/src/types"
)

func main() {
	// DataBase
	db.Connect()
	database := db.GetDB()
	database.AutoMigrate(&types.Lead{})
	log.Println("Database Migrated")

	// Server
	server := server.NewServer(fmt.Sprintf(":%s", "9010"), database)

	// RUN Server
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
