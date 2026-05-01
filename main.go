package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/waltertaya/cash-flow-forecast-backend/internals/api"
	"github.com/waltertaya/cash-flow-forecast-backend/internals/db"
	"github.com/waltertaya/cash-flow-forecast-backend/internals/middlewares"
)

func init() {
	db.Connect()

	// load the env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	api.SetupRoutes(r)

	r.Run(":8080")
}
