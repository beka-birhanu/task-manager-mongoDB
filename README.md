# Task Management Service

A task management service built with Go and MongoDB, providing functionality for adding, updating, deleting, and retrieving tasks. This service offers a RESTful API for managing tasks and demonstrates the use of MongoDB for data storage.

## Project Structure

The project is organized into the following directories:

- **`common`**: Contains common interfaces and error definitions.

  - `i_controller.go`: Interface for controllers.
  - `i_task_service.go`: Interface for task services.
  - `service_errors.go`: Defines common service errors.

- **`controllers`**: Contains the task controller and Data Transfer Objects (DTOs).

  - **`dto`**: Data Transfer Objects for request and response handling.
    - `add_request.go`: DTO for adding a task.
    - `get_response.go`: DTO for getting a task.
  - `task_controller.go`: Implements the task controller.

- **`data`**: Contains the implementation of the task service.

  - `task_service.go`: Implements the task management logic, including interactions with MongoDB.

- **`docs`**: Contains documentation for the API.

  - `api_definition.md`: API definition and usage documentation.

- **`models`**: Contains the data models used by the application.

  - `task.go`: Defines the Task model and its BSON representation.

- **`router`**: Contains the router configuration and setup.

  - `router.go`: Configures and initializes the router.

- **`main.go`**: Entry point for the application, initializing and running the server.

- **`README.md`**: Project documentation.

- **`go.mod`**: Go module file for dependency management.

- **`go.sum`**: Go module checksums for dependencies.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/beka-birhanu/task-manager-mongoDB.git
   ```

2. Change to the project directory:

   ```bash
   cd task-manager-mongoDB
   ```

3. Install the Go dependencies:
   ```bash
   go mod tidy
   ```

## Configuration

Before running the application, ensure you have a MongoDB instance running. Update the database connection string in main.go where it is set with clientOptions := options.Client().ApplyURI("<your-mongodb-connection-string>").

## Running the Application

To run the application, use:

```bash
go run main.go
```

The application will start a server on port `8080`. You can access the API at `http://localhost:8080/api/v1`.

## API Endpoints

- **Add Task**: `POST /api/v1/tasks`
- **Get All Tasks**: `GET /api/v1/tasks`
- **Get Task by ID**: `GET /api/v1/tasks/{id}`
- **Update Task**: `PUT /api/v1/tasks/{id}`
- **Delete Task**: `DELETE /api/v1/tasks/{id}`

Refer to `docs/api_definition.md` for detailed API usage and request/response formats.
