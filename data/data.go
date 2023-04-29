package data

import (
	"errors"
	"fmt"
	"gogin-api/models"

	"github.com/google/uuid"
)

type ToDoListDB struct {
	Lists map[string]*models.ToDoList
}

func NewToDoListDB(lists map[string]*models.ToDoList) *ToDoListDB {
	return &ToDoListDB{
		Lists: lists,
	}
}

func (db *ToDoListDB) CreateList(owner string, todos []*models.ToDo) (*models.ToDoList, error) {

	if db.Lists == nil {
		db.Lists = make(map[string]*models.ToDoList)
	}

	listKey := uuid.New().String()
	todosMap := make(map[string]*models.ToDo)

	for _, todo := range todos {
		todoKey := uuid.New().String()
		todo.Id = todoKey
		todo.ListId = listKey
		todo.Completed = false
		todosMap[todoKey] = todo
	}

	db.Lists[listKey] = &models.ToDoList{
		Id:    listKey,
		Owner: owner,
		Todos: todosMap,
	}

	return db.Lists[listKey], nil
}

func (db *ToDoListDB) CreateToDoInList(todo *models.ToDo, listId string) (*models.ToDo, error) {

	if _, hasList := db.Lists[listId]; hasList {

		todoKey := uuid.New().String()
		newToDo := &models.ToDo{
			Id:        todoKey,
			ListId:    listId,
			Content:   todo.Content,
			Completed: false,
		}
		db.Lists[listId].Todos[todoKey] = newToDo
		return newToDo, nil
	}

	return nil, errors.New("list not found")
}

func (db *ToDoListDB) GetList(listId string) (*models.ToDoList, error) {
	if _, hasList := db.Lists[listId]; hasList {
		return db.Lists[listId], nil

	}

	return nil, errors.New("list not found")
}

func (db *ToDoListDB) GetToDoInList(id string) (*models.ToDo, error) {
	for _, list := range db.Lists {
		if todo, hasToDo := list.Todos[id]; hasToDo {
			return &models.ToDo{
				ListId:    list.Id,
				Content:   todo.Content,
				Completed: todo.Completed,
			}, nil
		}
	}

	return nil, errors.New("todo not found")
}

func (db *ToDoListDB) GetAllLists() ([]models.ToDoList, string) {
	fmt.Printf("lists: %v", db.Lists)
	var lists []models.ToDoList

	for _, list := range db.Lists {
		lists = append(lists, *list)
	}

	if len(lists) == 0 {
		return nil, "no content"
	}

	return lists, ""

}

func (db *ToDoListDB) GetAllTodos(listId string) ([]*models.ToDo, error) {
	if _, hasList := db.Lists[listId]; hasList {
		var todos []*models.ToDo
		for _, todo := range db.Lists[listId].Todos {
			todos = append(todos, todo)
		}
		return todos, nil
	}

	return nil, errors.New("list not found")
}

func (db *ToDoListDB) PatchList(owner string, listId string) (*models.ToDoList, error) {
	if _, hasList := db.Lists[listId]; hasList {
		db.Lists[listId].Owner = owner
		return db.Lists[listId], nil
	}

	return nil, errors.New("list not found")
}

func (db *ToDoListDB) PatchToDoInList(isCompleted bool, todoId string) (*models.ToDo, error) {
	for _, list := range db.Lists {
		if todo, hasToDo := list.Todos[todoId]; hasToDo {
			if todo.Completed && isCompleted {
				return nil, errors.New("todo already completed")
			}

			if !todo.Completed && !isCompleted {
				return nil, errors.New("todo already not completed")
			}

			list.Todos[todoId].Completed = isCompleted
			return list.Todos[todoId], nil
		}
	}

	return nil, errors.New("todo not found")
}

func (db *ToDoListDB) DeleteList(listId string) error {
	if _, hasList := db.Lists[listId]; hasList {
		delete(db.Lists, listId)
		return nil
	}

	return errors.New("list not found")
}

func (db *ToDoListDB) DeleteToDoInList(todoId string) error {
	for k, list := range db.Lists {
		if _, hasToDo := list.Todos[todoId]; hasToDo {
			delete(db.Lists[k].Todos, todoId)
			return nil
		}
	}

	return errors.New("todo not found")
}

func (db *ToDoListDB) GetDataStructure() (map[string]*models.ToDoList, string) {
	if db.Lists == nil {
		return nil, "no content"
	}

	return db.Lists, ""
}
