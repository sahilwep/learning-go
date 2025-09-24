package service

import (
	"github.com/google/uuid"
	"github.com/sahilwep/learning-go/Go-Projects/Student-api/model"
	"github.com/sahilwep/learning-go/Go-Projects/Student-api/store"
)

// Student Service Constructor:
type StudentService struct {
	service *store.StudentStore
}

// Constructor for Student Service
func NewStudentService(stu *store.StudentStore) *StudentService {
	return &StudentService{service: stu}
}

// Create Student:
func (s *StudentService) CreateStudent(name string, address model.Addr, dob model.Dob) model.Student {

	student := model.Student{
		ID:      uuid.New().String(),
		Name:    name,
		Address: &address,
		Birth:   dob,
	}
	s.service.Add(student)
	return student
}

// Get Specific Student Details:
func (s *StudentService) GetStudent(id string) (model.Student, error) {
	return s.service.Get(id)
}

// Get All Student Details:
func (s *StudentService) List() []model.Student {
	return s.service.List()
}

// Update Student Details:
func (s *StudentService) Update(stu model.Student) error {
	return s.service.Update(stu)
}

// Delete Student Details:
func (s *StudentService) Delete(id string) error {
	return s.service.Delete(id)
}
