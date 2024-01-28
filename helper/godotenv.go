package helper

import (
	"log"

	"github.com/joho/godotenv"
)

func GoDoEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	return err
}
