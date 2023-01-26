package models

import (
	"errors"
	"time"
)

type Todo struct {
	ID        uint      `json:"id" gorm:"autoIncrement" gorm:"primaryKey: type:uuid"`
	Title     string    `json:"title" gorm:"unique"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UpdateTodoInput struct {
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
}

func (i UpdateTodoInput) Validate() error {
	if i.Title == nil && i.Completed == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
