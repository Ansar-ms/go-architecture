// server/server.go
package server

import (
	"go-architecture/pkg/config"
	"go-architecture/pkg/handler"
	"go-architecture/pkg/middleware"
	"go-architecture/pkg/repository"
	"go-architecture/pkg/services"

	"github.com/gin-gonic/gin"
)

func StartServer(cfg *config.Config) error {
	// Initialize Gin router
	router := gin.New()
	router.Use(middleware.Logger())
	router.Use(gin.Recovery())

	// Initialize repository, service, and handler
	userRepository := repository.NewUserRepository()
	userService := services.NewUserService(userRepository)

	// Register user routes
	handler.RegisterUserRoutes(router, userService)

	// Run server
	if err := router.Run(cfg.Server.Address); err != nil {
		return err
	}

	return nil
}
