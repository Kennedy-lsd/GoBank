package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	User string
	Name string
	SSL  string
}

func InitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loadig .env file")
	}

	return &Config{
		User: os.Getenv("USER"),
		Name: os.Getenv("NAME"),
		SSL:  os.Getenv("SSL"),
	}
}
