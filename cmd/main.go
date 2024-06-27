// cmd/main.go
package main

import (
	_ "go-architecture/cmd/docs"
	"go-architecture/pkg/config"
	"go-architecture/pkg/server"
	"log"
)

// @title User Management API
// @version 1.0
// @description This is a sample server for managing users.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization

// @schemes http https
func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Start the server
	if err := server.StartServer(cfg); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
