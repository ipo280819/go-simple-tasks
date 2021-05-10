package main

import (
	"fmt"
	"go-tasks/tasks"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RunServer(port string) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)

	loadTaskRoutes(router)

	fmt.Printf("Server listen on port %v \n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func loadTaskRoutes(router *mux.Router) {
	router.HandleFunc("/tasks", tasks.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", tasks.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", tasks.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", tasks.DeleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", tasks.UpdateTask).Methods("PUT")
}
