package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sahilwep/learning-go/Go-Projects/Student-api/model"
	"github.com/sahilwep/learning-go/Go-Projects/Student-api/service"
)

type StudentHandler struct {
	service *service.StudentService
}

func NewStudentHandler(service *service.StudentService) *StudentHandler {
	return &StudentHandler{service: service}
}

// Create Student:
func (s *StudentHandler) CreateStudent(c *gin.Context) {
	var stu model.Student

	if err := c.ShouldBindJSON(&stu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure Address is not nil before dereferencing
	var addr model.Addr
	if stu.Address != nil {
		addr = *stu.Address
	}

	created := s.service.CreateStudent(stu.Name, addr, stu.Birth)
	c.JSON(http.StatusCreated, created)
}

// Get By ID:
func (s *StudentHandler) GetStudent(c *gin.Context) {
	id := c.Param("id")

	stu, err := s.service.GetStudent(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stu)
}

// List ALL Available Student:
func (s *StudentHandler) ListStudents(c *gin.Context) {
	stu := s.service.List()

	// If it's empty -> return "Empty message"
	if len(stu) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"Empty": "nothing to display",
		})
		return
	}

	c.JSON(http.StatusOK, stu) // else return data
}

// Update: PUT req
func (s *StudentHandler) UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var stu model.Student // create empty instance of student

	// First bind incoming JSON with above empty student:
	if err := c.ShouldBindJSON(&stu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stu.ID = id // update the incoming book id here.

	if err := s.service.Update(stu); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stu) // else return sussufuly written
}

// Delete Student:
func (s *StudentHandler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	if err := s.service.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "student deleted successfully"})
}
