package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}
}

func PgnPath() string {
	path := os.Getenv("PGN_PATH")
	if path == "" {
		log.Fatal("PGN_PATH is not set. Please set it in your environment or in a .env file.")
	}
	return path
}
