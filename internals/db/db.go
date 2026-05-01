package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB is the shared database connection used across the application.
var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("database/cashflow.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Database connected successfully")
}
