package repository

import (
	"github.com/sinyavcev/todoGolang/pkg/repository/models"
	"gorm.io/gorm"
)

type Todo interface {
	Create(todo models.Todo) (uint, error)
	GetAll() ([]models.Todo, error)
	Update(todoId int, todo models.UpdateTodoInput) error
	Delete(todoId int) error
}

type Repository struct {
	Todo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Todo: NewTodoPostgres(db),
	}
}
