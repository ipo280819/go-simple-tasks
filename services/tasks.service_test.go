package services

import (
	"go-tasks/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockTaskRepository struct {
	mock.Mock
}

func (mock *mockTaskRepository) Save(task *entities.TaskDTO) (*entities.TaskDTO, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entities.TaskDTO), args.Error(1)
}

func (mock *mockTaskRepository) FindAll() ([]entities.TaskDTO, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entities.TaskDTO), args.Error(1)
}

func (mock *mockTaskRepository) Find(id string) (*entities.TaskDTO, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entities.TaskDTO), args.Error(1)
}

func (mock *mockTaskRepository) Delete(id string) (bool, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(bool), args.Error(1)
}

func (mock *mockTaskRepository) Update(id string, task *entities.TaskDTO) (bool, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(bool), args.Error(1)
}

func TestValidateEmptyTask(t *testing.T) {
	testService := NewTaskService(nil)

	err := testService.Validate(nil)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "Task is empty")
}

func TestValidateEmptyTaskName(t *testing.T) {
	testService := NewTaskService(nil)

	task := entities.TaskDTO{}
	err := testService.Validate(&task)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "Task name is empty")
}

func TestValidateTask(t *testing.T) {
	testService := NewTaskService(nil)

	task := entities.TaskDTO{Name: "Task 1"}
	err := testService.Validate(&task)
	assert.Nil(t, err)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(mockTaskRepository)

	task := entities.TaskDTO{ID: "identifier", Name: "Task name", Content: "Task desc"}

	mockRepo.On("FindAll").Return([]entities.TaskDTO{task}, nil)

	service := NewTaskService(mockRepo)

	result, _ := service.FindAll()

	mockRepo.AssertExpectations(t)

	assert.Equal(t, "identifier", result[0].ID)
	assert.Equal(t, "Task name", result[0].Name)
	assert.Equal(t, "Task desc", result[0].Content)
}
