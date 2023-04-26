package service

import (
	"errors"
	"gogin-api/data"
	"gogin-api/models"

	"github.com/google/uuid"
)

type ToDoListService struct {
	db data.ToDoListDB
}

func NewToDoListService(data data.ToDoListDB) *ToDoListService {
	return &ToDoListService{
		db: data,
	}
}

// create list nu returneaza o eroare pentru ca am acoperit toate scenariile care ar putea arunca ceva de genul. pe struct-ul de request body am binding required pe field-uri deci automat se apeleaza middleware-ul daca e sa nu fie ok, iar daca map-ul nu e initializat, se face aici. sper ca nu am ratat ceva. functia returneaza id-ul listei ca sa fie pus in location in header.

func (s *ToDoListService) CreateList(reqBody *models.RequestBodyList) string {

	if s.db.Lists == nil {
		s.db.Lists = make(map[string]*models.ToDoList)
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

	s.db.Lists[listKey] = &models.ToDoList{
		Id:    listKey,
		Owner: reqBody.Owner,
		Todos: todos,
	}

	return s.db.Lists[listKey].Id
}

func (s *ToDoListService) PatchList(list *models.ToDoList) error {
	if _, hasList := s.db.Lists[list.Id]; hasList {
		s.db.Lists[list.Id].Owner = list.Owner
		return nil
	} else {
		return errors.New("list not found")
	}
}

func (s *ToDoListService) GetList(id string) (models.ResponseBodyList, error) {
	if list, hasList := s.db.Lists[id]; hasList {
		todos, _ := s.GetAllToDosInList(id)

		return models.ResponseBodyList{
			Owner: list.Owner,
			Todos: todos,
		}, nil
	}
	return models.ResponseBodyList{}, errors.New("list not found")
}

func (s *ToDoListService) DeleteList(key string) error {
	if _, hasList := s.db.Lists[key]; hasList {
		delete(s.db.Lists, key)
		return nil
	}
	return errors.New("list not found")
}

func (s *ToDoListService) GetAllLists() ([]models.ResponseBodyList, error) {
	var lists []models.ResponseBodyList

	for _, list := range s.db.Lists {
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

func (s *ToDoListService) CreateToDoInList(listId string, content string) (string, error) {

	if _, hasList := s.db.Lists[listId]; hasList {
		key := uuid.New().String()
		s.db.Lists[listId].Todos[key] = &models.ToDo{
			Id:        key,
			ListId:    listId,
			Content:   content,
			Completed: false,
		}
		return s.db.Lists[listId].Todos[key].Id, nil
	}
	return "", errors.New("list not found")
}

func (s *ToDoListService) PatchToDoInList(completed bool, id string) error {
	for k, list := range s.db.Lists {
		if _, hasTodo := list.Todos[id]; hasTodo {
			s.db.Lists[k].Todos[id].Completed = completed
			return nil
		}
	}
	return errors.New("todo not found")
}

func (s *ToDoListService) GetToDoInList(key string) (*models.ToDo, error) {
	for _, list := range s.db.Lists {
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
	for k, list := range s.db.Lists {
		if _, hasToDo := list.Todos[key]; hasToDo {
			delete(s.db.Lists[k].Todos, key)
			return nil
		}
	}
	return errors.New("todo not found")
}

func (s *ToDoListService) GetAllToDosInList(listId string) ([]models.ToDo, error) {

	if _, hasList := s.db.Lists[listId]; hasList {
		var todos []models.ToDo
		for _, todo := range s.db.Lists[listId].Todos {
			responseToDo := models.ToDo{
				Id:        todo.Id,
				Content:   todo.Content,
				Completed: todo.Completed,
			}
			todos = append(todos, responseToDo)
		}
		return todos, nil
	}
	return nil, errors.New("list not found")

}

func (s *ToDoListService) GetDataStructure() (map[string]*models.ToDoList, error) {

	if s.db.Lists == nil {
		return nil, errors.New("no content")
	}

	return s.db.Lists, nil

}
