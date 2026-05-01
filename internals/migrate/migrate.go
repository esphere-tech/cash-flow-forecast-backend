package main

import (
	"github.com/waltertaya/cash-flow-forecast-backend/internals/db"
	"github.com/waltertaya/cash-flow-forecast-backend/internals/models"
)

func main() {
	db := db.Connect()

	db.AutoMigrate(&models.User{}, &models.CashEntry{})
}
