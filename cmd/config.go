package cmd

import (
	"os"
)

type Config struct {
	Port   string
	DBHost string
}

func GetConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbHost := os.Getenv("DB_URL")

	if dbHost == "" {
		dbHost = "localhost"
	}

	return Config{
		Port:   port,
		DBHost: dbHost,
	}
}
