package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sahilwep/learning-go/Go-Projects/Student-api/api"
	"github.com/sahilwep/learning-go/Go-Projects/Student-api/service"
	"github.com/sahilwep/learning-go/Go-Projects/Student-api/store"
)

func main() {
	r := gin.Default() // Initialize router.

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
