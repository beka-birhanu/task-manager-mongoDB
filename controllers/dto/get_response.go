package dto

import (
	"time"

	"github.com/beka-birhanu/models"
	"github.com/google/uuid"
)

type TaskResponse struct {
	ID          uuid.UUID     `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	DueDate     time.Time     `json:"dueDate"`
	Status      models.Status `json:"status"`
}
