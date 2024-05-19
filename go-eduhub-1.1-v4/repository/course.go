package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"
)

type CourseRepository interface {
	FetchByID(id int) (*model.Course, error)
	Store(course *model.Course) error
	ResetCourseRepo()
}

type courseRepository struct {
	courses []model.Course
}

func NewCourseRepo() *courseRepository {
	return &courseRepository{}
}

func (c *courseRepository) FetchByID(id int) (*model.Course, error) {
	for _, course := range c.courses {
		if course.ID == id {
			return &course, nil
		}
	}
	return nil, fmt.Errorf("course with id %d not found", id)
}

func (c *courseRepository) Store(course *model.Course) error {
	if course == nil {
		return fmt.Errorf("Invalid Add")
	}
	if course.ID == 0 {
		course.ID = len(c.courses) + 1
	}
	c.courses = append(c.courses, *course)
	return nil
}

func (c *courseRepository) ResetCourseRepo() {
	c.courses = []model.Course{}
}
