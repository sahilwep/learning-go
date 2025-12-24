package main

import (
	"auth-service-go/internal/config"
	"auth-service-go/internal/store"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load() // Loading configuration from /internal/config.go which will fetches all the env variables files..

	// Database Connection:
	db, err := store.NewPostgres(cfg)
	if err != nil {
		log.Fatal("DB connection Failed:", err)
	}
	defer db.Close()

	// Starting server:
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"db":     "connected",
		})
	})

	log.Println("Starting server on: ", cfg.AppPort)
	r.Run(":8080")

}
