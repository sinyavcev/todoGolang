package repository

import (
	"github.com/sinyavcev/todoGolang/repository/models"
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

//func (r *TodoListPostgres) GetById(userId, listId int) (todo.TodoList, error) {
//	var list todo.TodoList
//
//	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
//								INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
//		todoListsTable, usersListsTable)
//	err := r.db.Get(&list, query, userId, listId)
//
//	return list, err
//}
//
//func (r *TodoListPostgres) Delete(userId, listId int) error {
//	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
//		todoListsTable, usersListsTable)
//	_, err := r.db.Exec(query, userId, listId)
//
//	return err
//}
