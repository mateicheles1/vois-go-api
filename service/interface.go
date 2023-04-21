package service

import "gogin-api/models"

type ToDoListService interface {
	CreateList(reqBody *models.RequestBodyList)
	PatchList(list *models.ToDoList) error
	GetList(id string) (models.ResponseBodyList, error)
	DeleteList(key string) error
	GetAllLists() ([]models.ResponseBodyList, error)

	CreateToDoInList(listId string, content string) error
	PatchToDoInList(completed bool, id string) error
	GetToDoInList(key string) (*models.ToDo, error)
	DeleteToDoInList(key string) error
	GetAllToDosInList(listId string) ([]models.ToDo, error)
}
