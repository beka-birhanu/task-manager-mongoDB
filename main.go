package main

import (
	"log"

	"github.com/beka-birhanu/controllers"
	"github.com/beka-birhanu/data"
	"github.com/beka-birhanu/router"
)

// Configuration constants
const (
	addr    = ":8080"
	baseURL = "/api"
)

func main() {
	// Create a new task service instance
	taskService := data.NewTaskService()

	// Create a new task controller instance
	taskController := controllers.NewTaskController(taskService)

	// Create a new router instance with configuration
	routerConfig := router.Config{
		Addr:         addr,
		BaseURL:      baseURL,
		TasksHandler: taskController,
	}
	router := router.NewRouter(routerConfig)

	// Run the server
	if err := router.Run(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
