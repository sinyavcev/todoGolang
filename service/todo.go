package service

import (
	"github.com/sinyavcev/todoGolang/repository"
	"github.com/sinyavcev/todoGolang/repository/models"
)

type TodoService struct {
	repo repository.Todo
}

func NewTodoService(repo repository.Todo) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) Create(todo models.Todo) (uint, error) {
	return s.repo.Create(todo)
}

func (s *TodoService) GetAll() ([]models.Todo, error) {
	return s.repo.GetAll()
}

//func (s *TodoService) GetById(todoId int) (models.Todo, error) {
//	return s.repo.GetById(todoId)
//}
//
//func (s *TodoService) Delete(todoId int) error {
//	return s.repo.Delete(todoId)
//}
