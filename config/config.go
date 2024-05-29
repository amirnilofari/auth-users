package config

import "os"

type Config struct {
	DatabasePath string
	Port         string
}

func LoadConfig() Config {
	dbPath := os.Getenv("")
	if dbPath == "" {
		dbPath = "./data/myapp.db"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return Config{
		DatabasePath: dbPath,
		Port:         port,
	}
}
