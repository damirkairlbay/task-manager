package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Completed   bool      `gorm:"default:false" json:"completed"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("title cannot be empty")
	}

	if len(t.Title) > 255 {
		return errors.New("title is too long (max 255 characters)")
	}

	return nil
}

func (t *Task) BeforeCreate(tx *gorm.DB) error {
	return t.Validate()
}

func (t *Task) BeforeUpdate(tx *gorm.DB) error {
	return t.Validate()
}
