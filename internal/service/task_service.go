package service

import (
	"github.com/damirkairlbay/task-manager/internal/model"
	"github.com/damirkairlbay/task-manager/internal/repository"
)

type TaskService interface {
	CreateTask(title, description string, completed bool) (*model.Task, error)
	ListTasks() ([]model.Task, error)
	GetTask(id int64) (*model.Task, error)
	UpdateTask(id int64, title, description string, completed bool) (*model.Task, error)
	DeleteTask(id int64) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(title, description string, completed bool) (*model.Task, error) {
	task := &model.Task{
		Title:       title,
		Description: description,
		Completed:   completed,
	}

	if err := task.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Create(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *taskService) ListTasks() ([]model.Task, error) {
	return s.repo.List()
}

func (s *taskService) GetTask(id int64) (*model.Task, error) {
	return s.repo.GetByID(id)
}

func (s *taskService) UpdateTask(id int64, title, description string, completed bool) (*model.Task, error) {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	task.Title = title
	task.Description = description
	task.Completed = completed

	// Валидация изменений
	if err := task.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Update(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *taskService) DeleteTask(id int64) error {
	return s.repo.Delete(id)
}
