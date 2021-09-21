package repositories

import (
	"context"
	"go-tasks/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func newMongoTaskRepository() TaskRepository {
	return &mongoTaskRepository{}
}

type mongoTaskRepository struct {
}

var taskCollection, _ = GetMongoDBCollection("tasks")

func (*mongoTaskRepository) Save(task *entities.TaskDTO) (*entities.TaskDTO, error) {
	taskMongo := task.ToMongoTask()
	res, err := taskCollection.InsertOne(context.Background(), taskMongo)
	if err != nil {
		return nil, err
	}
	task.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return task, nil
}
func (*mongoTaskRepository) FindAll() ([]entities.TaskDTO, error) {
	ctx := context.Background()
	cursor, err := taskCollection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}
	var tasks []entities.TaskDTO
	for cursor.Next(ctx) {
		var taskMongo entities.TaskMongo
		err := cursor.Decode(&taskMongo)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, taskMongo.ToDTOTask())
	}
	return tasks, nil
}
func (*mongoTaskRepository) Find(id string) (*entities.TaskDTO, error) {
	return nil, nil
}
func (*mongoTaskRepository) Delete(id string) (bool, error) {
	return false, nil
}
func (*mongoTaskRepository) Update(id string, task *entities.TaskDTO) (bool, error) {
	return false, nil
}
