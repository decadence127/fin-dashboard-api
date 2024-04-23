package main

import (
	"log"
	"os"

	mongoInstance "fin-dashboard-api/app"
	"fin-dashboard-api/app/routes"

	"github.com/lpernett/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	if err := mongoInstance.ConnectToMongoDB(); err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	router := routes.GetRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router.Run(":" + port)
}
