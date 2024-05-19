package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	filebased *filebased.Data
}

func NewTaskRepo(filebasedDb *filebased.Data) *taskRepository {
	return &taskRepository{
		filebased: filebasedDb,
	}
}

func (t *taskRepository) Store(task *model.Task) error {
	t.filebased.StoreTask(*task)

	return nil
}

func (t *taskRepository) Update(task *model.Task) error {
	return t.filebased.UpdateTask(task.ID, *task)
}

func (t *taskRepository) Delete(id int) error {
	return t.filebased.DeleteTask(id)
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	return t.filebased.GetTaskByID(id)
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	return t.filebased.GetTasks()
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	return t.filebased.GetTaskListByCategory(id)
}
