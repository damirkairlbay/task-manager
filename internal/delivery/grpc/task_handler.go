package grpc

import (
	"context"

	"github.com/damirkairlbay/task-manager/internal/service"
	"github.com/damirkairlbay/task-manager/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TaskHandler struct {
	proto.UnimplementedTaskServiceServer
	service service.TaskService
}

func NewTaskHandler(service service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(ctx context.Context, req *proto.CreateTaskRequest) (*proto.TaskResponse, error) {
	task, err := h.service.CreateTask(req.Title, req.Description, req.Completed)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to create task: %v", err)
	}

	return &proto.TaskResponse{
		Task: &proto.Task{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
			CreatedAt:   task.CreatedAt.Unix(),
			UpdatedAt:   task.UpdatedAt.Unix(),
		},
	}, nil
}

func (h *TaskHandler) ListTasks(ctx context.Context, req *proto.ListTasksRequest) (*proto.ListTasksResponse, error) {
	tasks, err := h.service.ListTasks()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list tasks: %v", err)
	}

	var protoTasks []*proto.Task
	for _, task := range tasks {
		protoTasks = append(protoTasks, &proto.Task{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
			CreatedAt:   task.CreatedAt.Unix(),
			UpdatedAt:   task.UpdatedAt.Unix(),
		})
	}

	return &proto.ListTasksResponse{
		Tasks: protoTasks,
	}, nil
}

// GetTask возвращает задачу по ID через gRPC
func (h *TaskHandler) GetTask(ctx context.Context, req *proto.GetTaskRequest) (*proto.TaskResponse, error) {
	task, err := h.service.GetTask(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get task: %v", err)
	}

	return &proto.TaskResponse{
		Task: &proto.Task{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
			CreatedAt:   task.CreatedAt.Unix(),
			UpdatedAt:   task.UpdatedAt.Unix(),
		},
	}, nil
}

func (h *TaskHandler) UpdateTask(ctx context.Context, req *proto.UpdateTaskRequest) (*proto.TaskResponse, error) {
	task, err := h.service.UpdateTask(req.Id, req.Title, req.Description, req.Completed)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to update task: %v", err)
	}

	return &proto.TaskResponse{
		Task: &proto.Task{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
			CreatedAt:   task.CreatedAt.Unix(),
			UpdatedAt:   task.UpdatedAt.Unix(),
		},
	}, nil
}

func (h *TaskHandler) DeleteTask(ctx context.Context, req *proto.DeleteTaskRequest) (*proto.DeleteTaskResponse, error) {
	err := h.service.DeleteTask(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to delete task: %v", err)
	}

	return &proto.DeleteTaskResponse{
		Success: true,
	}, nil
}
