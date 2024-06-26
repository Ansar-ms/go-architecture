// cmd/main.go
package main

import (
	"go-architecture/pkg/config"
	"go-architecture/pkg/server"
	"log"
)

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
