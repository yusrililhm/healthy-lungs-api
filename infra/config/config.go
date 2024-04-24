package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type appConfig struct {
	AppPort    string
	DBUser     string
	DBHost     string
	DBPort     string
	DBName     string
	DBPassword string
	DBDialect  string
}

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("error while load .env: %s\n", err.Error())
	}
}

func AppConfig() *appConfig {
	return &appConfig{
		AppPort:    os.Getenv("APP_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBDialect:  os.Getenv("DB_DIALECT"),
	}
}
