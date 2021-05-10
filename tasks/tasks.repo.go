package tasks

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type TaskRepository interface {
	Save(task *Task) (*Task, error)
	FindAll() ([]Task, error)
}

type firebaseRepository struct {
}

func NewTaskRepository() TaskRepository {
	return &firebaseRepository{}
}

const (
	projectId      = "go-simple-tasks"
	collectionName = "Tasks"
)

func (*firebaseRepository) Save(task *Task) (*Task, error) {

	ctx := context.Background()
	client, err := CreateFirestoreClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	dr, _, err := client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"name":    task.Name,
		"content": task.Content,
	})

	if err != nil {
		fmt.Println("Error creating task in firestore:", err)
		return nil, err
	}
	task.ID = dr.ID
	return task, nil
}

func (*firebaseRepository) FindAll() ([]Task, error) {
	ctx := context.Background()
	client, err := CreateFirestoreClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	var tasks []Task
	iter := client.Collection(collectionName).Documents(ctx)
	var task Task
	for {

		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("Error getting tasks from firestore:", err)
			return nil, err
		}
		if err := doc.DataTo(&task); err != nil {

			fmt.Println("Error parsing tasks data from firestore:", err)
			return nil, err
		}

		task.ID = doc.Ref.ID
		tasks = append(tasks, task)

	}
	return tasks, nil
}

func CreateFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		fmt.Println("Error creating firestore client", err)
		return nil, err
	}
	return client, nil

}
