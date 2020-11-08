package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config (key string) returns a key from .env file
func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Errorf("Error loading .env file\n")
	}
	return os.Getenv(key)
}
