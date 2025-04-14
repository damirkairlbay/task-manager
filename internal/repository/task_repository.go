package repository

import (
	"errors"

	"github.com/damirkairlbay/task-manager/internal/model"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *model.Task) error
	List() ([]model.Task, error)
	GetByID(id int64) (*model.Task, error)
	Update(task *model.Task) error
	Delete(id int64) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) List() ([]model.Task, error) {
	var tasks []model.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) GetByID(id int64) (*model.Task, error) {
	var task model.Task
	if err := r.db.First(&task, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) Update(task *model.Task) error {
	if task.ID == 0 {
		return errors.New("task ID is required")
	}

	result := r.db.Save(task)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (r *taskRepository) Delete(id int64) error {
	result := r.db.Delete(&model.Task{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("task not found")
	}

	return nil
}
