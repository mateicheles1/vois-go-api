package service

import (
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

func (s *ToDoListService) GetLists() ([]*models.ToDoList, error) {

	lists, err := s.db.GetLists()

	for i := range lists {
		for j := range lists[i].Todos {
			lists[i].Todos[j].ListId = ""
		}
	}

	if err != nil {
		return nil, err
	}

	return lists, nil
}

func (s *ToDoListService) GetTodos(listId string) ([]*models.ToDo, error) {
	todos, err := s.db.GetTodos(listId)

	if err != nil {
		return nil, err
	}

	for i := range todos {
		todos[i].ListId = ""
	}

	return todos, nil
}

func (s *ToDoListService) CreateList(requestBody *models.RequestBodyList) (*models.ToDoList, error) {

	var todos []*models.ToDo

	for _, todo := range requestBody.Todos {
		dbToDo := &models.ToDo{
			ListId:    requestBody.Id,
			Content:   todo,
			Completed: false,
		}

		todos = append(todos, dbToDo)
	}

	requestBody.Todos = nil

	createdList, err := s.db.CreateList(requestBody, todos)

	if err != nil {
		return nil, err
	}

	return createdList, nil

}

func (s *ToDoListService) GetList(listId string) (*models.ToDoList, error) {
	list, err := s.db.GetList(listId)

	list.Id = ""

	for i, todo := range list.Todos {
		list.Todos[i] = &models.ToDo{
			Id:        todo.Id,
			Content:   todo.Content,
			Completed: todo.Completed,
		}
	}

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s *ToDoListService) PatchList(reqBody *models.RequestBodyList, listId string) (*models.ToDoList, error) {

	list, err := s.db.PatchList(reqBody, listId)

	for i, v := range list.Todos {
		list.Todos[i] = &models.ToDo{
			Id:        v.Id,
			Content:   v.Content,
			Completed: v.Completed,
		}
	}

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s *ToDoListService) DeleteList(listId string) error {
	err := s.db.DeleteList(listId)

	if err != nil {
		return err
	}

	return nil
}

func (s *ToDoListService) CreateTodo(reqBody *models.ToDo, listId string) (*models.ToDo, error) {

	todo, err := s.db.CreateTodo(reqBody, listId)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *ToDoListService) GetTodo(todoId string) (*models.ToDo, error) {
	todo, err := s.db.GetTodo(todoId)

	if err != nil {
		return nil, err
	}

	responseTodo := &models.ToDo{
		ListId:    todo.ListId,
		Content:   todo.Content,
		Completed: todo.Completed,
	}

	return responseTodo, nil
}

func (s *ToDoListService) PatchTodo(reqBody *models.ToDo, todoId string) (*models.ToDo, error) {
	todo, err := s.db.PatchTodo(reqBody, todoId)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *ToDoListService) DeleteTodo(todoId string) error {
	err := s.db.DeleteTodo(todoId)

	if err != nil {
		return err
	}

	return nil
}
