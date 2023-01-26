package service

import (
	"github.com/sinyavcev/todoGolang/pkg/repository"
	"github.com/sinyavcev/todoGolang/pkg/repository/models"
)

type Todo interface {
	Create(todo models.Todo) (uint, error)
	GetAll() ([]models.Todo, error)
	Update(todoId int, todo models.UpdateTodoInput) error
	Delete(todoId int) error
}

type Service struct {
	Todo
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Todo: NewTodoService(repos.Todo),
	}
}
