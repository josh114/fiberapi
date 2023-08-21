package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB 
}

var Database DbInstance

func connectDB () {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database", err.Error())
		os.Exit(2 )
	}
	log.Println("Connected to the database successfully")
	db.Logger = logger.Dafault.LogMode(logger.Info)
	log.Println("Running migrations")
}