package main

import (
	"log"
	"pt-xyz/internal/delivery/server"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, using system environment variables.")
	}

	if err := server.Run(); err != nil {
		log.Fatalf("Failed to run server: %e", err)
	}
}
