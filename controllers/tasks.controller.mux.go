package controllers

import (
	"encoding/json"
	"fmt"
	"go-tasks/entities"
	"net/http"

	"github.com/gorilla/mux"
)

type muxController struct {
}

func NewTaskMuxController() TaskController {
	return &muxController{}
}

func (*muxController) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := taskService.FindAll()
	if err != nil {
		responseError(w, err)
		return
	}
	responseOK(w, tasks)
}

func (*muxController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask entities.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		responseError(w, err)
		return
	}

	result, err := taskService.Create(&newTask)

	if err != nil {
		responseError(w, err)
		return
	}
	responseStatus(http.StatusCreated, w, result)
}

func (*muxController) GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	task, err := taskService.Find(id)
	if err != nil {
		if err.Error() == "Not Found" {
			responseErrorStatus(http.StatusNotFound, w, err)
			return
		}
		responseError(w, err)
		return
	}
	responseOK(w, task)
}
func (*muxController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_ = vars["id"]
}

func (*muxController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_ = vars["id"]
}

func responseOK(w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
func responseStatus(status int, w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(result)
}

func responseError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Error: %v", err)
}
func responseErrorStatus(status int, w http.ResponseWriter, err error) {
	w.WriteHeader(status)
	fmt.Fprintf(w, "Error: %v", err)
}
