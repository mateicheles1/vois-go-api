package data

import "gogin-api/models"

type ToDoListDB struct {
	Lists map[string]*models.ToDoList
}
