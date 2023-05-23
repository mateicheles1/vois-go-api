package service

import "gogin-api/models"

type ToDoListServiceInterface interface {
	CreateList(requestBody *models.RequestBodyList, owner string) (*models.ToDoList, error)
	CreateTodo(reqBody *models.ToDo, listId string) (*models.ToDo, error)
	CreateUser(reqBody *models.User) (*models.User, error)
	DeleteList(listId string) error
	DeleteAllLists(role string) error
	DeleteTodo(todoId string) error
	DeleteUser(username string, role string) error
	GetList(listId string, username string, role string) (*models.ToDoList, error)
	GetLists(username string, role string) ([]*models.ToDoList, error)
	GetTodo(todoId string) (*models.ToDo, error)
	GetTodos(listId string) ([]*models.ToDo, error)
	PatchList(reqBody *models.RequestBodyList, listId string) (*models.ToDoList, error)
	PatchTodo(reqBody *models.ToDo, todoId string) (*models.ToDo, error)
	Login(reqBody *models.User) (string, error)
}
