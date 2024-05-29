package main

import (
	"database/sql"

	"github.com/amirnilofari/auth-users/config"
	"github.com/amirnilofari/auth-users/internal/cli"
	"github.com/amirnilofari/auth-users/internal/repository"
	"github.com/amirnilofari/auth-users/internal/service"
	"github.com/gofiber/fiber/v2/log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := config.LoadConfig()

	db, err := sql.Open("sqlite3", cfg.DatabasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Run database migration
	if _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            first_name TEXT NOT NULL,
			last_name TEXT NOT NULL,
            email TEXT NOT NULL UNIQUE
        );
    `); err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	cli.StartCLI(userService)
}
