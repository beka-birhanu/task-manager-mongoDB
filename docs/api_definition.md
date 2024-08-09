# Task Management API

This API allows users to manage tasks. It provides endpoints to create, update, delete, and retrieve tasks.
[postman docs](https://documenter.getpostman.com/view/37481680/2sA3s3Fqkb)
## Endpoints

### Create a Task

**URL:** `/tasks`

**Method:** `POST`

**Request Body:**

```json
{
  "title": "string",
  "description": "string",
  "dueDate": "string (ISO 8601 format)",
  "status": "string"
}
```

**Response:**

- **Status:** `201 Created`
- **Headers:**
  - `Location`: The URL of the created task.

**Errors:**

- `400 Bad Request`: Invalid request data.

### Update a Task

**URL:** `/tasks/{id}`

**Method:** `PUT`

**Path Parameters:**

- `id` (UUID): The ID of the task to update.

**Request Body:**

```json
{
  "title": "string",
  "description": "string",
  "dueDate": "string (ISO 8601 format)",
  "status": "string"
}
```

**Response:**

- **Status:** `200 OK`

**Errors:**

- `400 Bad Request`: Invalid request data.
- `404 Not Found`: Task not found.

### Delete a Task

**URL:** `/tasks/{id}`

**Method:** `DELETE`

**Path Parameters:**

- `id` (UUID): The ID of the task to delete.

**Response:**

- **Status:** `200 OK`

**Errors:**

- `400 Bad Request`: Invalid task ID.
- `404 Not Found`: Task not found.

### Get All Tasks

**URL:** `/tasks`

**Method:** `GET`

**Response:**

- **Status:** `200 OK`
- **Body:**

```json
[
  {
    "id": "uuid",
    "title": "string",
    "description": "string",
    "dueDate": "string (ISO 8601 format)",
    "status": "string",
    "createdAt": "string (ISO 8601 format)",
    "updatedAt": "string (ISO 8601 format)"
  }
]
```

### Get a Single Task

**URL:** `/tasks/{id}`

**Method:** `GET`

**Path Parameters:**

- `id` (UUID): The ID of the task to retrieve.

**Response:**

- **Status:** `200 OK`
- **Body:**

```json
{
  "id": "uuid",
  "title": "string",
  "description": "string",
  "dueDate": "string (ISO 8601 format)",
  "status": "string",
  "createdAt": "string (ISO 8601 format)",
  "updatedAt": "string (ISO 8601 format)"
}
```

**Errors:**

- `400 Bad Request`: Invalid task ID.
- `404 Not Found`: Task not found.

## Models

### AddTaskRequest

```json
{
  "title": "string",
  "description": "string",
  "dueDate": "string (ISO 8601 format)",
  "status": "string"
}
```

### UpdateTaskRequest

```json
{
  "title": "string",
  "description": "string",
  "dueDate": "string (ISO 8601 format)",
  "status": "string"
}
```

### TaskResponse

```json
{
  "id": "uuid",
  "title": "string",
  "description": "string",
  "dueDate": "string (ISO 8601 format)",
  "status": "string",
  "createdAt": "string (ISO 8601 format)",
  "updatedAt": "string (ISO 8601 format)"
}
```
