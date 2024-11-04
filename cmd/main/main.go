package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mohanapranes/book-store/pkg/routes"
	"github.com/mohanapranes/book-store/pkg/utils"
)

func main() {
	// Set up Gin router with default middleware (logger and recovery)
	router := gin.Default()

	// Register application routes
	routes.RegisterRoutes(router)

	// Define server address using environment variable, defaulting to localhost:9010
	addr := utils.GetEnv("SERVER_ADDRESS", "localhost:9010")

	// Configure the HTTP server
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	// Start the server in a goroutine
	go func() {
		log.Printf("Starting server on %s", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")

	// Gracefully shut down with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
