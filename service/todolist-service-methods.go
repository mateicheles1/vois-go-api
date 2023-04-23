package service

import (
	"errors"
	"gogin-api/models"

	"github.com/google/uuid"
)

func (s *ToDoListService) CreateList(reqBody *models.RequestBodyList) {
	if s.repo.Lists == nil {
		s.repo.Lists = make(map[string]*models.ToDoList)
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

	s.repo.Lists[listKey] = &models.ToDoList{
		Id:    listKey,
		Owner: reqBody.Owner,
		Todos: todos,
	}
}

func (s *ToDoListService) PatchList(list *models.ToDoList) error {
	if _, hasList := s.repo.Lists[list.Id]; hasList {
		s.repo.Lists[list.Id].Owner = list.Owner
		return nil
	} else {
		return errors.New("list not found")
	}
}

func (s ToDoListService) GetList(id string) (models.ResponseBodyList, error) {
	if list, hasList := s.repo.Lists[id]; hasList {
		todos, _ := s.GetAllToDosInList(id)

		return models.ResponseBodyList{
			Owner: list.Owner,
			Todos: todos,
		}, nil
	}
	return models.ResponseBodyList{}, errors.New("list not found")
}

func (s *ToDoListService) DeleteList(key string) error {
	if _, hasList := s.repo.Lists[key]; hasList {
		delete(s.repo.Lists, key)
		return nil
	}
	return errors.New("list not found")
}

func (s ToDoListService) GetAllLists() ([]models.ResponseBodyList, error) {
	var lists []models.ResponseBodyList

	for _, list := range s.repo.Lists {
		todos, _ := s.GetAllToDosInList(list.Id)
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

func (s *ToDoListService) CreateToDoInList(listId string, content string) error {
	if _, hasList := s.repo.Lists[listId]; hasList {
		key := uuid.New().String()
		s.repo.Lists[listId].Todos[key] = &models.ToDo{
			Id:        key,
			ListId:    listId,
			Content:   content,
			Completed: false,
		}
		return nil
	}
	return errors.New("list not found")
}

func (s *ToDoListService) PatchToDoInList(completed bool, id string) error {
	for k, list := range s.repo.Lists {
		if _, hasTodo := list.Todos[id]; hasTodo {
			s.repo.Lists[k].Todos[id].Completed = completed
			return nil
		}
	}
	return errors.New("tood not found")
}

func (s *ToDoListService) GetToDoInList(key string) (*models.ToDo, error) {
	for _, list := range s.repo.Lists {
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
func (s *ToDoListService) DeleteToDoInList(key string) error {
	for k, list := range s.repo.Lists {
		if _, hasToDo := list.Todos[key]; hasToDo {
			delete(s.repo.Lists[k].Todos, key)
			return nil
		}
	}
	return errors.New("todo not found")
}

func (s *ToDoListService) GetAllToDosInList(listId string) ([]models.ToDo, error) {
	if _, hasList := s.repo.Lists[listId]; hasList {
		var todos []models.ToDo
		for _, todo := range s.repo.Lists[listId].Todos {
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

func (s ToDoListService) GetDataStructure() (map[string]*models.ToDoList, error) {
	if s.repo.Lists == nil {
		return nil, errors.New("no content")
	} else {
		return s.repo.Lists, nil
	}
}
