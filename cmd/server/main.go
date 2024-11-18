package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/flowkater/ddd-todo-app/config"
	"github.com/flowkater/ddd-todo-app/wire"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Initialize server using Wire
	server, err := wire.InitializeServer(cfg)
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}

	// Start server
	go func() {
		if err := server.Start(); err != nil {
			log.Printf("server error: %v\n", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down server...")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("server shutdown error: %v\n", err)
	}
}
