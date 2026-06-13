package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
}

func Load() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	cfg := &Config{
		ServerPort: os.Getenv("SERVER_PORT"),
	}

	return cfg, nil
}
