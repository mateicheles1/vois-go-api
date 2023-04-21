package data

import "gogin-api/models"

type ToDoListRepo struct {
	Lists map[string]*models.ToDoList
}
