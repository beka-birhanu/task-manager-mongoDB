package data

import (
	"time"

	"github.com/beka-birhanu/common"
	"github.com/beka-birhanu/models"
	"github.com/google/uuid"
)

// TaskService represents a service for managing tasks.
type TaskService struct {
	store map[uuid.UUID]*models.Task
}

// NewTaskService creates a new TaskService.
func NewTaskService() *TaskService {
	return &TaskService{
		store: make(map[uuid.UUID]*models.Task),
	}
}

// Add adds a new task to the store. Returns an error if there is an ID conflict.
func (s *TaskService) Add(title, description string, dueDate time.Time, status models.Status) (*models.Task, error) {
	taskConfig := models.TaskConfig{
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Status:      status,
	}

	// Recreate task unitil the id conflict is resolved
	for {
		task, err := models.NewTask(taskConfig)
		if err != nil {
			return &models.Task{}, err
		}

		if _, exists := s.store[task.ID()]; exists {
			continue
		}
		s.store[task.ID()] = task
		return task, nil
	}
}

// Update updates an existing task. Returns an error if the task is not found.
func (s *TaskService) Update(id uuid.UUID, title, description string, dueDate time.Time, status models.Status) (*models.Task, error) {
	task, exists := s.store[id]
	if !exists {
		return &models.Task{}, common.ErrNotFound
	}

	taskConfig := models.TaskConfig{
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Status:      status,
	}

	if err := task.Update(taskConfig); err != nil {
		return task, err
	}
	return task, nil
}

// Delete removes a task by ID. Returns an error if the task is not found.
func (s *TaskService) Delete(id uuid.UUID) error {
	if _, exists := s.store[id]; !exists {
		return common.ErrNotFound
	}
	delete(s.store, id)
	return nil
}

// GetAll returns a list of pointers to all tasks.
func (s *TaskService) GetAll() []*models.Task {
	tasks := make([]*models.Task, 0, len(s.store))
	for _, task := range s.store {
		tasks = append(tasks, task)
	}
	return tasks
}

// GetSingle returns a task by ID. Returns an error if the task is not found.
func (s *TaskService) GetSingle(id uuid.UUID) (*models.Task, error) {
	task, exists := s.store[id]
	if !exists {
		return nil, common.ErrNotFound
	}
	return task, nil
}
