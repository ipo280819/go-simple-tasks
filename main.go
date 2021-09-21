package main

import (
	"go-tasks/constants"
	"go-tasks/controllers"
	router "go-tasks/http"
	"go-tasks/repositories"
	"go-tasks/services"
	"log"

	"github.com/joho/godotenv"
)

const (
	useRouter     = constants.MUX
	useRepository = constants.MONGO
)

var (
	httpRouter router.Router = router.NewRouter(useRouter)
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	loadTaskRoutes()
	httpRouter.SERVE(":3000")
}

var (
	taskRepository repositories.TaskRepository = repositories.NewTaskRepository(useRepository)
	taskService    services.TaskService        = services.NewTaskService(taskRepository)
	taskController controllers.TaskController  = controllers.NewTaskController(taskService, useRouter)
)

func loadTaskRoutes() {
	httpRouter.GET("/tasks", taskController.GetTasks)
	httpRouter.POST("/tasks", taskController.CreateTask)
	httpRouter.GET("/tasks/{id}", taskController.GetTask)
	httpRouter.DELETE("/tasks/{id}", taskController.DeleteTask)
	httpRouter.PUT("/tasks/{id}", taskController.UpdateTask)
}
