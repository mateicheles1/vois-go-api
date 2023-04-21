package models

type ToDoListService interface {
	CreateList(reqBody *RequestBodyList)
	PatchList(list *ToDoList) error
	GetList(id string) (ResponseBodyList, error)
	DeleteList(key string) error
	GetAllLists() ([]ResponseBodyList, error)

	CreateToDoInList(listId string, content string) error
	PatchToDoInList(completed bool, id string) error
	GetToDoInList(key string) (*ToDo, error)
	DeleteToDoInList(key string) error
	GetAllToDosInList(listId string) ([]ToDo, error)
}
