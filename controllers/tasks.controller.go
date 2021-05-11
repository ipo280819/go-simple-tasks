package controllers

import (
	"go-tasks/constants"
	"go-tasks/services"
	"net/http"
)

var (
	taskService services.TaskService
)

type TaskController interface {
	GetTasks(w http.ResponseWriter, r *http.Request)
	CreateTask(w http.ResponseWriter, r *http.Request)
	GetTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
}

func NewTaskController(service services.TaskService, typeRouter string) TaskController {
	taskService = service

	switch typeRouter {
	case constants.MUX:
		return NewTaskMuxController()
	}
	return nil
}
