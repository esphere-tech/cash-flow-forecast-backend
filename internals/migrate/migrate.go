package main

import (
	"github.com/waltertaya/cash-flow-forecast-backend/internals/db"
	"github.com/waltertaya/cash-flow-forecast-backend/internals/models"
)

func main() {
	db.Connect()

	db.DB.AutoMigrate(&models.User{}, &models.CashEntry{})
}
