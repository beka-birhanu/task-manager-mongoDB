package common

import (
	"time"

	"github.com/beka-birhanu/models"
	"github.com/google/uuid"
)

type ITaskService interface {

	// Add adds a new task to the store. Returns an error if there is an ID conflict.
	Add(title, description string, dueDate time.Time, status models.Status) (*models.Task, error)

	// Update updates an existing task. Returns an error if the task is not found.
	Update(id uuid.UUID, title, description string, dueDate time.Time, status models.Status) (*models.Task, error)

	// Delete removes a task by ID. Returns an error if the task is not found.
	Delete(id uuid.UUID) error

	// GetAll returns a list of pointers to all tasks.
	GetAll() []*models.Task

	// GetSingle returns a task by ID. Returns an error if the task is not found.
	GetSingle(id uuid.UUID) (*models.Task, error)
}
