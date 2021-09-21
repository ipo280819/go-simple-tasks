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
	var newTask entities.TaskDTO
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
	id := vars["id"]
	wasDeleted, err := taskService.Delete(id)
	if err != nil {
		if err.Error() == "Not Found" {
			responseErrorStatus(http.StatusNotFound, w, err)
			return
		}
		responseError(w, err)
		return
	}

	result := TaskDeletedDTO{
		id,
		wasDeleted,
	}
	responseOK(w, result)
}

func (*muxController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var task entities.TaskDTO
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		responseError(w, err)
		return
	}
	wasUpdated, err := taskService.Update(id, &task)
	if err != nil {
		if err.Error() == "Not Found" {
			responseErrorStatus(http.StatusNotFound, w, err)
			return
		}
		responseError(w, err)
		return
	}

	result := TaskUpdatedDTO{task, wasUpdated}
	responseOK(w, result)
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
