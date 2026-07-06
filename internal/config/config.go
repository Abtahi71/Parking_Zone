package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	Dsn       string
	JwtSecret string
}

func LoadEnv() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	return &Config{
		Port: os.Getenv("Port"),
		Dsn: os.Getenv("Dsn"),
		JwtSecret: os.Getenv("JwtSecret"),
	}
}