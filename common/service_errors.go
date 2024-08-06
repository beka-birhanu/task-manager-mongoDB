package common

import "errors"

// ErrIDConflict represents an error when a task with the same ID already exists.
var ErrIDConflict = errors.New("task ID conflict")

// ErrNotFound represents an error when a task is not found.
var ErrNotFound = errors.New("task not found")
