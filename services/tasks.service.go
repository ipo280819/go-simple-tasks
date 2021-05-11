package services

import (
	"errors"
	"go-tasks/entities"
	"go-tasks/repositories"
)

type TaskService interface {
	Validate(task *entities.Task) error
	Create(task *entities.Task) (*entities.Task, error)
	FindAll() ([]entities.Task, error)
	Find(id string) (*entities.Task, error)
}

type service struct {
}

var (
	repo repositories.TaskRepository
)

func NewTaskService(repository repositories.TaskRepository) TaskService {
	repo = repository
	return &service{}
}

func (*service) Validate(task *entities.Task) error {
	if task == nil {
		err := errors.New("Task is empty")
		return err
	}
	if task.Name == "" {
		err := errors.New("Task name is empty")
		return err
	}
	return nil
}

func (this *service) Create(task *entities.Task) (*entities.Task, error) {
	err := this.Validate(task)
	if err != nil {
		return nil, err
	}
	return repo.Save(task)
}

func (*service) FindAll() ([]entities.Task, error) {
	return repo.FindAll()
}

func (*service) Find(id string) (*entities.Task, error) {
	return repo.Find(id)
}
