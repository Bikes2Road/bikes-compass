package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Bikes2Road/bikes-compass/cmd/api/config"
	"github.com/Bikes2Road/bikes-compass/cmd/api/wrapper"
	_ "github.com/Bikes2Road/bikes-compass/docs"
)

// @title           Bikes Compass API
// @version         1.0
// @description     This is the docs of Bikes Compass API from Bikes2Road.

// @BasePath  /api/v1/bikes

func main() {

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	log.Printf("Starting Bikes Compass Microservice in %s mode...", cfg.Server.Env)

	w := wrapper.DefaultWrapper()

	// Initialize dependency injection container
	app, err := wrapper.NewApp(w, cfg)
	if err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}

	// Setup router
	router := app.Router.SetUp(cfg.Server.IsDevelopment())

	// Create HTTP server
	srv := &http.Server{
		Addr:         cfg.Server.GetBindAddress(),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server listening on %s", cfg.Server.GetBindAddress())
		log.Printf("Swagger documentation available at http://%s/api/v1/bikes/swagger/index.html", cfg.Server.GetServerAddress())
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited successfully")

}
