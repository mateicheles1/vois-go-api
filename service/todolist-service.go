package service

import "gogin-api/data"

type ToDoListService struct {
	repo data.ToDoListRepo
}

func NewToDoListService(repo data.ToDoListRepo) *ToDoListService {
	return &ToDoListService{
		repo: repo,
	}
}
