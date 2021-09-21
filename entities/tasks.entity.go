package entities

import (
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskDTO struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type TaskMongo struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `json:"name"`
	Content string             `json:"content"`
}

func (_task *TaskMongo) ToDTOTask() TaskDTO {
	var taskDTO TaskDTO
	copier.Copy(&taskDTO, &_task)
	if !_task.ID.IsZero() {
		taskDTO.ID = _task.ID.Hex()
	}
	return taskDTO
}
func (_taskDTO *TaskDTO) ToMongoTask() TaskMongo {
	var task TaskMongo
	copier.Copy(&task, &_taskDTO)
	id, err := primitive.ObjectIDFromHex(_taskDTO.ID)
	if err != nil {
		task.ID = id
	}
	return task
}
