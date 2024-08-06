package dto

import (
	"time"

	"github.com/beka-birhanu/models"
)

// DTOs for task operations
type AddTaskRequest struct {
	Title       string        `json:"title" binding:"required"`
	Description string        `json:"description" binding:"required"`
	DueDate     time.Time     `json:"dueDate" binding:"required"`
	Status      models.Status `json:"status" binding:"required"`
}
