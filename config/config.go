package config

import (
	"fmt"
	"os"
)

type Config struct {
	DatabasePath string
	Port         string
}

func LoadConfig() Config {
	os.Setenv("DB_PATH", "./users.db")
	dbPath := os.Getenv("DB_PATH")
	fmt.Println("db:", os.Getenv("DB_PATH"))
	if dbPath == "" {
		dbPath = "./data/myapp.db"
	}
	os.Unsetenv("DB_PATH")

	os.Setenv("PORT", "3000")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	os.Unsetenv("PORT")

	return Config{
		DatabasePath: dbPath,
		Port:         port,
	}
}
