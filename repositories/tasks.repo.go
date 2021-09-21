package repositories

import (
	"go-tasks/constants"
	"go-tasks/entities"
)

type TaskRepository interface {
	Save(task *entities.TaskDTO) (*entities.TaskDTO, error)
	FindAll() ([]entities.TaskDTO, error)
	Find(id string) (*entities.TaskDTO, error)
	Delete(id string) (bool, error)
	Update(id string, task *entities.TaskDTO) (bool, error)
}

func NewTaskRepository(typeRepo string) TaskRepository {

	switch typeRepo {
	case constants.FIRESTORE:
		return newFirestoreTaskRepository()

	case constants.MONGO:
		return newMongoTaskRepository()
	}
	return nil
}
