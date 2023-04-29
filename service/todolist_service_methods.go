package service

import (
	"errors"
	"gogin-api/data"
	"gogin-api/models"
)

type ToDoListService struct {
	db data.ToDoListDBInterface
}

func NewToDoListService(data data.ToDoListDBInterface) *ToDoListService {
	return &ToDoListService{
		db: data,
	}
}

func (s *ToDoListService) CreateList(requestBody models.RequestBodyList) (*models.ToDoList, error) {
	if requestBody.Owner == "" {
		return nil, errors.New("empty owner")
	}
	var todos []*models.ToDo
	for _, stringContent := range requestBody.Todos {
		newToDo := &models.ToDo{
			Content: stringContent,
		}

		todos = append(todos, newToDo)
	}

	list, err := s.db.CreateList(requestBody.Owner, todos)

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s *ToDoListService) CreateToDoInList(requestBody *models.ToDo, listId string) (*models.ToDo, error) {

	if requestBody.Content == "" {
		return nil, errors.New("empty content")
	}

	todo, err := s.db.CreateToDoInList(requestBody, listId)

	if err != nil {
		return nil, err
	}

	return todo, nil

}

func (s *ToDoListService) GetList(listId string) (*models.ResponseBodyList, error) {
	list, err := s.db.GetList(listId)
	var todos []*models.ToDo

	if err != nil {
		return nil, err
	}

	for _, todo := range list.Todos {
		newToDo := &models.ToDo{
			Id:        todo.Id,
			Content:   todo.Content,
			Completed: todo.Completed,
		}

		todos = append(todos, newToDo)
	}

	return &models.ResponseBodyList{
		Owner: list.Owner,
		Todos: todos,
	}, nil
}

func (s *ToDoListService) GetToDoInList(id string) (*models.ToDo, error) {
	todo, err := s.db.GetToDoInList(id)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *ToDoListService) GetAllLists() ([]models.ResponseBodyList, string) {

	var responseLists []models.ResponseBodyList

	lists, message := s.db.GetAllLists()

	for _, list := range lists {
		todos, _ := s.GetAllTodos(list.Id)
		reponseList := models.ResponseBodyList{
			Id:    list.Id,
			Owner: list.Owner,
			Todos: todos,
		}

		responseLists = append(responseLists, reponseList)
	}

	if message != "" {
		return nil, message
	}

	return responseLists, ""

}

func (s *ToDoListService) GetAllTodos(listId string) ([]*models.ToDo, error) {
	todos, err := s.db.GetAllTodos(listId)

	var responseTodos []*models.ToDo

	if err != nil {
		return nil, err
	}

	for _, todo := range todos {
		responseToDo := &models.ToDo{
			Id:        todo.Id,
			Content:   todo.Content,
			Completed: todo.Completed,
		}

		responseTodos = append(responseTodos, responseToDo)
	}

	return responseTodos, nil

}

func (s *ToDoListService) PatchList(requestBody *models.ToDoList, listId string) (*models.ToDoList, error) {

	if requestBody.Owner == "" {
		return nil, errors.New("invalid JSON syntax in request body; empty owner")
	}

	list, err := s.db.PatchList(requestBody.Owner, listId)

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s *ToDoListService) PatchToDoInList(requestBody *models.ToDo, todoId string) (*models.ToDo, error) {
	todo, err := s.db.PatchToDoInList(requestBody.Completed, todoId)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *ToDoListService) DeleteList(listId string) error {
	err := s.db.DeleteList(listId)

	if err != nil {
		return err
	}

	return nil
}

func (s *ToDoListService) DeleteToDoInList(todoId string) error {
	err := s.db.DeleteToDoInList(todoId)

	if err != nil {
		return err
	}

	return nil
}

func (s *ToDoListService) GetDataStructure() (map[string]*models.ToDoList, string) {
	data, message := s.db.GetDataStructure()

	if message != "" {
		return nil, message
	}

	return data, ""
}
