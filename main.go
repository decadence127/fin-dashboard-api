package main

import (
	"log"
	"os"

	"github.com/lpernett/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	if err := connectToMongoDB(); err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	router := getRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router.Run(":" + port)
}
