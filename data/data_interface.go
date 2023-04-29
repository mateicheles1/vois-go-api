package data

import "gogin-api/models"

type ToDoListDBInterface interface {
	CreateList(owner string, todos []*models.ToDo) (*models.ToDoList, error)
	CreateToDoInList(todo *models.ToDo, listId string) (*models.ToDo, error)
	GetList(listId string) (*models.ToDoList, error)
	GetToDoInList(id string) (*models.ToDo, error)
	GetAllLists() ([]models.ToDoList, string)
	GetAllTodos(listId string) ([]*models.ToDo, error)
	GetDataStructure() (map[string]*models.ToDoList, string)
	PatchList(owner string, listId string) (*models.ToDoList, error)
	PatchToDoInList(isCompleted bool, todoId string) (*models.ToDo, error)
	DeleteList(listId string) error
	DeleteToDoInList(todoId string) error
}
