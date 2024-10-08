package main

import (
	"context"
	"log"

	"github.com/beka-birhanu/controllers"
	"github.com/beka-birhanu/data"
	"github.com/beka-birhanu/router"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Configuration constants
const (
	addr    = ":8080"
	baseURL = "/api"
)

func main() {
	// Initialize MongoDB client
	clientOptions := options.Client().ApplyURI("")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %v", err)
	}

	// Ping the MongoDB server to verify connection
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("Error pinging MongoDB server: %v", err)
	}

	// Create a new task service instance
	taskService := data.NewTaskService(client, "taskdb", "tasks")

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

