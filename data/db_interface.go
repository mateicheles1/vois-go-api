package data

import "gogin-api/models"

type ToDoListDBInterface interface {
	CreateList(reqBody *models.RequestBodyList, todos []*models.ToDo) (*models.ToDoList, error)
	CreateTodo(reqBody *models.ToDo, listId string) (*models.ToDo, error)
	CreateUser(reqBody *models.User) (*models.User, error)
	DeleteList(listId string) error
	DeleteAllLists() error
	DeleteTodo(todoId string) error
	GetAllListsAdmin() ([]*models.ToDoList, error)
	GetList(listId string) (*models.ToDoList, error)
	GetLists(username string) ([]*models.ToDoList, error)
	GetTodo(todoId string) (*models.ToDo, error)
	GetTodos(listId string) ([]*models.ToDo, error)
	PatchList(reqBody *models.RequestBodyList, listId string) (*models.ToDoList, error)
	PatchTodo(reqBody *models.ToDo, todoId string) (*models.ToDo, error)
	Login(reqBody *models.User) (*models.User, error)
}
