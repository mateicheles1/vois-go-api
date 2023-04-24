package service

import "gogin-api/data"

type ToDoListService struct {
	db data.ToDoListDB
}

func NewToDoListService(data data.ToDoListDB) *ToDoListService {
	return &ToDoListService{
		db: data,
	}
}
