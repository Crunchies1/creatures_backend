package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Crunchies1/creatures_backend/handlers"
	"github.com/Crunchies1/creatures_backend/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbClient, err := setupMongoConnection()
	if err != nil {
		log.Fatal(err)
	}
	modelsClient := models.NewClient(dbClient, os.Getenv("MONGO_DB_NAME"))

	router := gin.Default()
	app := handlers.New(router, modelsClient)
	app.SetupRoutes(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Graceful shutdown setup
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// Shutdown HTTP server
		log.Println("🔽 Shutting down HTTP server...")
		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}

		// Disconnect Mongo client
		if err := dbClient.Disconnect(context.Background()); err != nil {
			log.Printf("MongoDB Disconnect error: %v", err)
		} else {
			log.Println("✅ MongoDB disconnected")
		}

		close(idleConnsClosed)
	}()

	log.Println("🚀 Server listening on :8080")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
	log.Println("👋 Server exited")
}
