package repository

import (
	"github.com/sinyavcev/todoGolang/pkg/repository/models"
	"gorm.io/gorm"
)

type TodoPostgres struct {
	db *gorm.DB
}

func NewTodoPostgres(db *gorm.DB) *TodoPostgres {
	return &TodoPostgres{db: db}
}

func (r *TodoPostgres) Create(todo models.Todo) (uint, error) {
	result := r.db.Create(&todo) // pass pointer of data to Create
	return todo.ID, result.Error
}

func (r *TodoPostgres) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	result := r.db.Find(&todos)
	return todos, result.Error
}

func (r *TodoPostgres) Delete(todoId int) error {
	result := r.db.Delete(&models.Todo{}, todoId)
	return result.Error
}

func (r *TodoPostgres) Update(todoId int, input models.UpdateTodoInput) error {
	res := r.db.Model(models.Todo{}).Where("id = ?", todoId).Updates(&input)
	return res.Error
}
