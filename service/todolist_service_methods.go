package service

import (
	"errors"
	"gogin-api/data"
	"gogin-api/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type ToDoListService struct {
	db data.ToDoListDBInterface
}

func NewToDoListService(data data.ToDoListDBInterface) *ToDoListService {
	return &ToDoListService{
		db: data,
	}
}

var invalidUUID = UuidError{message: "invalid uuid format"}

func (s *ToDoListService) GetLists(username string, role string) ([]*models.ToDoList, error) {

	var lists []*models.ToDoList
	var err error

	if role != "admin" {
		lists, err = s.db.GetLists(username)
	} else {
		lists, err = s.db.GetAllListsAdmin()
	}

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

	if _, err := uuid.Parse(listId); err != nil {
		return nil, invalidUUID
	}

	todos, err := s.db.GetTodos(listId)

	if err != nil {
		return nil, err
	}

	for i := range todos {
		todos[i].ListId = ""
	}

	return todos, nil
}

func (s *ToDoListService) CreateList(requestBody *models.RequestBodyList, owner string) (*models.ToDoList, error) {

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

func (s *ToDoListService) GetList(listId string, username string, role string) (*models.ToDoList, error) {

	if _, err := uuid.Parse(listId); err != nil {
		return nil, invalidUUID
	}

	list, err := s.db.GetList(listId)

	if err != nil {
		return nil, err
	}

	if username != list.Owner || role != "admin" {
		return nil, errors.New("action not allowed")
	}

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

	if _, err := uuid.Parse(listId); err != nil {
		return nil, invalidUUID
	}

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

	if _, err := uuid.Parse(listId); err != nil {
		return invalidUUID
	}

	err := s.db.DeleteList(listId)

	if err != nil {
		return err
	}

	return nil
}

func (s *ToDoListService) DeleteAllLists(role string) error {

	if role != "admin" {
		return errors.New("action not allowed")
	}

	err := s.db.DeleteAllLists()

	if err != nil {
		return err
	}

	return nil
}

func (s *ToDoListService) CreateTodo(reqBody *models.ToDo, listId string) (*models.ToDo, error) {

	if _, err := uuid.Parse(listId); err != nil {
		return nil, invalidUUID
	}

	todo, err := s.db.CreateTodo(reqBody, listId)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *ToDoListService) GetTodo(todoId string) (*models.ToDo, error) {

	if _, err := uuid.Parse(todoId); err != nil {
		return nil, invalidUUID
	}

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

	if _, err := uuid.Parse(todoId); err != nil {
		return nil, invalidUUID
	}

	todo, err := s.db.PatchTodo(reqBody, todoId)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *ToDoListService) DeleteTodo(todoId string) error {

	if _, err := uuid.Parse(todoId); err != nil {
		return invalidUUID
	}

	err := s.db.DeleteTodo(todoId)

	if err != nil {
		return err
	}

	return nil
}

func (s *ToDoListService) CreateUser(reqBody *models.User) (*models.User, error) {

	user, err := s.db.CreateUser(reqBody)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *ToDoListService) DeleteUser(username string, role string) error {

	if role != "admin" {
		return errors.New("action not allowed")
	}

	err := s.db.DeleteUser(username)

	if err != nil {
		return err
	}

	return nil
}

func (s *ToDoListService) Login(reqBody *models.User) (string, error) {

	user, err := s.db.Login(reqBody)

	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
