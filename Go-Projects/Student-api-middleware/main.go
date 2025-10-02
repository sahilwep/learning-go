package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sahilwep/learning-go/Go-Projects/Student-api/api"
	"github.com/sahilwep/learning-go/Go-Projects/Student-api/middleware"
	"github.com/sahilwep/learning-go/Go-Projects/Student-api/service"
	"github.com/sahilwep/learning-go/Go-Projects/Student-api/store"
)

func main() {

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load")
	} else {
		log.Println(".env file loaded successfully")
	}

	r := gin.Default() // Initialize router.

	// Middleware order: recovery -> logger -> CORS -> rate limit -> timeout
	r.Use(
		middleware.RecoveryMiddleware(),                    // recovery middleware
		middleware.Logger(true),                            // logger middleware -> {flag = true: we need our logs into files}
		middleware.CORS(),                                  // handel cross origin
		middleware.RateLimiterMiddleware(5, 1*time.Minute), // 5 requests per minutes per IP
		middleware.SimpleTimeoutMiddleware(8*time.Second),  // Timeout Middleware
	)

	// Create pieces of applications:
	stu := store.NewStudentStore()
	stuService := service.NewStudentService(stu)
	handler := api.NewStudentHandler(stuService)

	// Register routes:
	r.POST("/student", handler.CreateStudent)
	r.GET("/student", handler.ListStudents)
	r.GET("/student/:id", handler.GetStudent)
	r.PUT("/student/:id", handler.UpdateStudent)
	r.DELETE("/student/:id", handler.DeleteStudent)

	// Serve:
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
