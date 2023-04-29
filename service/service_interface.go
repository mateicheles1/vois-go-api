package service

import "gogin-api/models"

type ToDoListServiceInterface interface {
	CreateList(requestBody models.RequestBodyList) (*models.ToDoList, error)
	CreateToDoInList(requestBody *models.ToDo, listId string) (*models.ToDo, error)
	DeleteList(listId string) error
	DeleteToDoInList(todoId string) error
	GetList(listId string) (*models.ResponseBodyList, error)
	GetToDoInList(id string) (*models.ToDo, error)
	GetAllLists() ([]models.ResponseBodyList, string)
	GetAllTodos(listId string) ([]*models.ToDo, error)
	GetDataStructure() (map[string]*models.ToDoList, string)
	PatchList(requestBody *models.ToDoList, listId string) (*models.ToDoList, error)
	PatchToDoInList(requestBody *models.ToDo, todoId string) (*models.ToDo, error)
}
