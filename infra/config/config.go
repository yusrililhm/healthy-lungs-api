package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv()  {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("error while load .env: %s\n", err.Error()) 
	}
}
