package store

import (
	"errors"

	"github.com/sahilwep/learning-go/Go-Projects/Student-api/model"
)

var ErrorNotFound = errors.New("student not found")

type StudentStore struct {
	students map[string]model.Student
}

// Constructor for StudentStore:
func NewStudentStore() *StudentStore {
	return &StudentStore{
		students: make(map[string]model.Student),
	}
}

// Create Method:
func (s *StudentStore) Add(stu model.Student) {
	s.students[stu.ID] = stu
}

// Get by ID Method:
func (s *StudentStore) Get(id string) (model.Student, error) {
	stu, found := s.students[id]

	if !found {
		return model.Student{}, ErrorNotFound // if not found return: empty student{} & error
	}

	return stu, nil
}

// List Method:
func (s *StudentStore) List() []model.Student {
	all := []model.Student{}

	for _, val := range s.students {
		all = append(all, val)
	}

	return all
}

// Update Method:
func (s *StudentStore) Update(stu model.Student) error {
	_, ok := s.students[stu.ID]

	if !ok {
		return ErrorNotFound
	}

	s.students[stu.ID] = stu
	return nil
}

// Delete Method:
func (s *StudentStore) Delete(id string) error {
	_, ok := s.students[id]

	if !ok {
		return ErrorNotFound
	}

	delete(s.students, id) // delete(map, key) to delete any {key:val} from map
	return nil
}
