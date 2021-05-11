package repositories

import (
	"context"
	"errors"
	"fmt"
	"go-tasks/entities"
	"go-tasks/utils"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func newFirestoreTaskRepository() TaskRepository { return &firebaseRepository{} }

type firebaseRepository struct {
}

const (
	projectId      = "go-simple-tasks"
	collectionName = "Tasks"
)

func (*firebaseRepository) Save(task *entities.Task) (*entities.Task, error) {

	ctx := context.Background()
	client, err := CreateFirestoreClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	taskCreate := utils.StructToMap(task)
	dr, _, err := client.Collection(collectionName).Add(ctx, taskCreate)

	if err != nil {
		fmt.Println("Error creating task in firestore:", err)
		return nil, err
	}
	task.ID = dr.ID
	return task, nil
}

func (*firebaseRepository) FindAll() ([]entities.Task, error) {
	ctx := context.Background()
	client, err := CreateFirestoreClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	var tasks []entities.Task
	iter := client.Collection(collectionName).Documents(ctx)
	var task entities.Task
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
func (*firebaseRepository) Find(id string) (*entities.Task, error) {

	ctx := context.Background()
	client, err := CreateFirestoreClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	var task entities.Task
	doc, err := client.Collection(collectionName).Doc(id).Get(ctx)
	if !doc.Exists() {
		err = errors.New("Not Found")
		return nil, err
	}
	if err := doc.DataTo(&task); err != nil {

		fmt.Println("Error parsing tasks data from firestore:", err)
		return nil, err
	}
	task.ID = id
	return &task, nil
}
func (*firebaseRepository) Delete(id string) (bool, error) {

	ctx := context.Background()
	client, err := CreateFirestoreClient(ctx)
	if err != nil {
		return false, err
	}
	defer client.Close()
	_, err = client.Collection(collectionName).Doc(id).Delete(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (*firebaseRepository) Update(id string, task *entities.Task) (bool, error) {

	ctx := context.Background()
	client, err := CreateFirestoreClient(ctx)
	if err != nil {
		return false, err
	}
	defer client.Close()
	taskUpdate := utils.StructToMap(task)
	_, err = client.Collection(collectionName).Doc(id).Set(ctx, taskUpdate)
	if err != nil {
		return false, err
	}
	return true, nil
}

func CreateFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		fmt.Println("Error creating firestore client", err)
		return nil, err
	}
	return client, nil

}
