package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	database, err := gorm.Open(sqlite.Open("crm.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect Database")
	} else {
		log.Println("DataBase Succecfully Connected")
	}

	db = database
}

func GetDB() *gorm.DB {
	return db
}
