package service

import (
	"github.com/sinyavcev/todoGolang/repository"
	"github.com/sinyavcev/todoGolang/repository/models"
)

type Todo interface {
	Create(todo models.Todo) (uint, error)
	GetAll() ([]models.Todo, error)
	//GetById(todoId int) (models.Todo, error)
	//Delete(todoId int) error
}

type Service struct {
	Todo
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Todo: NewTodoService(repos.Todo),
	}
}
