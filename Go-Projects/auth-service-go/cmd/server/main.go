package main

import (
	"auth-service-go/internal/config"
	"auth-service-go/internal/router"
	"auth-service-go/internal/store"
	"auth-service-go/internal/user"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load Config:
	cfg := config.Load()

	// Connect DB
	db, err := store.NewPostgres(cfg)
	if err != nil {
		log.Fatal("DB connection Failed:", err)
	}
	defer db.Close()

	// Init repository
	userRepo := user.NewRepository(db)

	// Start Gin
	r := gin.Default()

	// Health Check:
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"db":     "connected",
		})
	})

	// Auth Routes:
	router.RegisterAuthRoutes(r, userRepo)

	log.Println("Starting server on: ", cfg.AppPort)
	r.Run(":" + cfg.AppPort)
}
