package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/waltertaya/cash-flow-forecast-backend/internals/api"
	"github.com/waltertaya/cash-flow-forecast-backend/internals/db"
	"github.com/waltertaya/cash-flow-forecast-backend/internals/middlewares"
)

func init() {
	// load the env variables
	if err := godotenv.Load(); err != nil {
		log.Printf(".env file not found, using environment variables")
	}

	db.Connect()
}

func main() {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	api.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
