# Task Manager gRPC Service

A service for task management using gRPC, Go, and PostgreSQL.

## Project Description

This service implements basic CRUD operations for task management (Task Manager) through a gRPC API.

### Main operations:

- `CreateTask` - create a new task
- `ListTasks` - retrieve a list of all tasks
- `GetTask` - get a task by ID
- `UpdateTask` - update an existing task
- `DeleteTask` - delete a task by ID

## Technologies

- Go 1.20+
- gRPC
- PostgreSQL
- GORM (ORM for Go)
- Docker and Docker Compose

## Architecture

The project follows a layered architecture:

- **Model** - defines data structures and business logic
- **Repository** - data access layer
- **Service** - business logic layer
- **Delivery** - gRPC handlers

## Installation and Setup

### Prerequisites

- Docker and Docker Compose
- Go 1.20+ (for local development)
- protoc (if you need to generate proto files)

### Running with Docker Compose

1. Clone the repository:
```bash
git clone https://github.com/yourusername/task-manager.git
cd task-manager
```

2. Run with Docker Compose:
```bash
docker-compose up -d
```

The service will be available on port 50051.

### Running Locally

1. Start PostgreSQL:
```bash
docker-compose up -d db
```

2. Install dependencies:
```bash
go mod download
```

3. Run the server:
```bash
go run cmd/server/main.go
```

## Using the gRPC API

### Request Examples (using grpcurl)

Install grpcurl:
```bash
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

#### Create a task

```bash
grpcurl -d '{"title": "New Task", "description": "Task description", "completed": false}' \
  -plaintext localhost:50051 proto.TaskService/CreateTask
```

#### Get the list of tasks

```bash
grpcurl -plaintext localhost:50051 proto.TaskService/ListTasks
```

#### Get a task by ID

```bash
grpcurl -d '{"id": 1}' -plaintext localhost:50051 proto.TaskService/GetTask
```

#### Update a task

```bash
grpcurl -d '{"id": 1, "title": "Updated Task", "description": "New description", "completed": true}' \
  -plaintext localhost:50051 proto.TaskService/UpdateTask
```

#### Delete a task

```bash
grpcurl -d '{"id": 1}' -plaintext localhost:50051 proto.TaskService/DeleteTask
```

## Generating Proto Files

If you need to modify the gRPC protocols and regenerate Go files:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  pkg/proto/task.proto
```

## Error Handling

The service implements error handling:
- Data validation at the model level
- Detailed error messages
- Appropriate gRPC status codes

## License

MIT