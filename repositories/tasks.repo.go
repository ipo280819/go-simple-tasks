package repositories

import (
	"go-tasks/constants"
	"go-tasks/entities"
)

type TaskRepository interface {
	Save(task *entities.Task) (*entities.Task, error)
	FindAll() ([]entities.Task, error)
	Find(id string) (*entities.Task, error)
	Delete(id string) (bool, error)
	Update(id string, task *entities.Task) (bool, error)
}

func NewTaskRepository(typeRepo string) TaskRepository {

	switch typeRepo {
	case constants.FIRESTORE:
		return newFirestoreTaskRepository()
	}
	return nil
}
