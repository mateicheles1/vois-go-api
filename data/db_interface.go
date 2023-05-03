package data

import "gogin-api/models"

type ToDoListDBInterface interface {
	CreateList(reqBody *models.RequestBodyList, todos []*models.ToDo) (*models.ToDoList, error)
	GetList(listId string) (*models.ToDoList, error)
	GetLists() ([]*models.ToDoList, error)
	GetTodos(listId string) ([]*models.ToDo, error)
	PatchList(reqBody *models.RequestBodyList, listId string) (*models.ToDoList, error)
	DeleteList(listId string) error
	CreateTodo(reqBody *models.ToDo, listId string) (*models.ToDo, error)
	GetTodo(todoId string) (*models.ToDo, error)
	PatchTodo(reqBody *models.ToDo, todoId string) (*models.ToDo, error)
	DeleteTodo(todoId string) error
}
