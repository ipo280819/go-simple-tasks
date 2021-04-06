package handlers

import (
	"encoding/json"
	"fmt"
	"go-tasks/my_types"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var tasks = my_types.AllTasks{
	{
		ID:      1,
		Name:    "Task 1",
		Content: "Content 1",
	},
	{
		ID:      2,
		Name:    "Task 2",
		Content: "Content 2",
	},
	{
		ID:      3,
		Name:    "Task 3",
		Content: "Content 3",
	},
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprint(w, "Error: Invalid ID")
		return
	}
	for _, task := range tasks {
		if task.ID == id {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusFound)
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask my_types.Task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprint(w, "Error: Invalid ID")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)

			response := make(map[string]interface{})
			response["tasks"] = tasks
			response["taskDeleted"] = task

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusFound)
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprint(w, "Error: Invalid ID")
		return
	}
	var taskUpdated my_types.Task

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	json.Unmarshal(reqBody, &taskUpdated)
	taskUpdated.ID = id
	for i, task := range tasks {
		if task.ID == id {

			tasks[i] = taskUpdated
			response := make(map[string]interface{})
			response["tasks"] = tasks
			response["taskUpdated"] = taskUpdated

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusFound)
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
