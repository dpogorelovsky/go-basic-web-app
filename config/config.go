package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig loads .env config from root of the project to current environment
func LoadConfig() {
	log.Println("loading .env config")
	godotenv.Load()
}

// Get - config getter
func Get(name string) string {
	return os.Getenv(name)
}
