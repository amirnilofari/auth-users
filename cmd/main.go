package main

import (
	"github.com/amirnilofari/auth-users/internal/cli"
	"github.com/amirnilofari/auth-users/internal/repository"
	"github.com/amirnilofari/auth-users/internal/service"
)

func main() {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	cli.StartCLI(userService)
}
