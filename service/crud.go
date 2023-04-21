package service

import (
	"errors"
	"gogin-api/models"

	"github.com/google/uuid"
)

type ToDoListRepo struct {
	Lists map[string]*models.ToDoList
}

func (r *ToDoListRepo) CreateList(reqBody *models.RequestBodyList) {
	if r.Lists == nil {
		r.Lists = make(map[string]*models.ToDoList)
	}
	listKey := uuid.New().String()

	todos := make(map[string]*models.ToDo)

	for _, v := range reqBody.Todos {
		toDosKey := uuid.New().String()
		todos[toDosKey] = &models.ToDo{
			Id:        toDosKey,
			ListId:    listKey,
			Content:   v,
			Completed: false,
		}
	}

	r.Lists[listKey] = &models.ToDoList{
		Id:    listKey,
		Owner: reqBody.Owner,
		Todos: todos,
	}
}

func (r *ToDoListRepo) PatchList(list *models.ToDoList) error {
	if _, hasList := r.Lists[list.Id]; hasList {
		r.Lists[list.Id].Owner = list.Owner
		return nil
	} else {
		return errors.New("list not found")
	}
}

func (r ToDoListRepo) GetList(id string) (models.ResponseBodyList, error) {
	if list, hasList := r.Lists[id]; hasList {
		todos, _ := r.GetAllToDosInList(id)

		return models.ResponseBodyList{
			Owner: list.Owner,
			Todos: todos,
		}, nil
	}
	return models.ResponseBodyList{}, errors.New("list not found")
}

func (r *ToDoListRepo) DeleteList(key string) error {
	if _, hasList := r.Lists[key]; hasList {
		delete(r.Lists, key)
		return nil
	}
	return errors.New("list not found")
}

func (r ToDoListRepo) GetAllLists() ([]models.ResponseBodyList, error) {
	var lists []models.ResponseBodyList

	for _, list := range r.Lists {
		todos, _ := r.GetAllToDosInList(list.Id)
		responseList := models.ResponseBodyList{
			Id:    list.Id,
			Owner: list.Owner,
			Todos: todos,
		}
		lists = append(lists, responseList)
	}
	if len(lists) == 0 {
		return nil, errors.New("no content")
	}
	return lists, nil
}

func (r *ToDoListRepo) CreateToDoInList(listId string, content string) error {
	if _, hasList := r.Lists[listId]; hasList {
		key := uuid.New().String()
		r.Lists[listId].Todos[key] = &models.ToDo{
			Id:        key,
			ListId:    listId,
			Content:   content,
			Completed: false,
		}
		return nil
	}
	return errors.New("list not found")
}

func (r *ToDoListRepo) PatchToDoInList(completed bool, id string) error {
	for k, list := range r.Lists {
		if _, hasTodo := list.Todos[id]; hasTodo {
			r.Lists[k].Todos[id].Completed = completed
			return nil
		}
	}
	return errors.New("tood not found")
}

func (r *ToDoListRepo) GetToDoInList(key string) (*models.ToDo, error) {
	for _, list := range r.Lists {
		if todo, hasToDo := list.Todos[key]; hasToDo {
			return &models.ToDo{
				ListId:    list.Id,
				Content:   todo.Content,
				Completed: todo.Completed,
			}, nil
		}
	}
	return nil, errors.New("todo not found")
}
func (r *ToDoListRepo) DeleteToDoInList(key string) error {
	for k, list := range r.Lists {
		if _, hasToDo := list.Todos[key]; hasToDo {
			delete(r.Lists[k].Todos, key)
			return nil
		}
	}
	return errors.New("todo not found")
}

func (r *ToDoListRepo) GetAllToDosInList(listId string) ([]models.ToDo, error) {
	if _, hasList := r.Lists[listId]; hasList {
		var todos []models.ToDo
		for _, todo := range r.Lists[listId].Todos {
			responseToDo := &models.ToDo{
				Id:        todo.Id,
				Content:   todo.Content,
				Completed: todo.Completed,
			}
			todos = append(todos, *responseToDo)
		}
		return todos, nil
	}
	return nil, errors.New("list not found")

}
