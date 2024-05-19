package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	Store(student *model.Student) error
	FindByID(id int) (student model.Student, err error)
	ResetStudentRepo()
}

type studentRepository struct {
	students []model.Student
}

func NewStudentRepo() *studentRepository {
	return &studentRepository{}
}

func (s *studentRepository) FetchAll() ([]model.Student, error) {
	return s.students, nil
}

func (s *studentRepository) FindByID(id int) (student model.Student, err error) {
	isFound := false
	for _, v := range s.students {
		if v.ID == id {
			student = v
			isFound = true
			break
		}
	}
	if isFound {
		return student, nil
	}
	return student, fmt.Errorf("student with id %d not found", id)
}

func (s *studentRepository) Store(student *model.Student) error {
	if student == nil {
		return fmt.Errorf("Invalid Add")
	}

	if student.ID == 0 {
		student.ID = len(s.students) + 1
	}

	s.students = append(s.students, *student)

	return nil
}

func (s *studentRepository) ResetStudentRepo() {
	s.students = []model.Student{}
}
