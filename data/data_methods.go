package data

import (
	"errors"
	"gogin-api/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ToDoListDB struct {
	lists *gorm.DB
}

func NewToDoListDB(db *gorm.DB) *ToDoListDB {
	return &ToDoListDB{
		lists: db,
	}
}

func (db *ToDoListDB) CreateList(reqBody *models.RequestBodyList, todos []*models.ToDo) (*models.ToDoList, error) {

	listId := uuid.New().String()

	dbList := models.ToDoList{
		Id:    listId,
		Owner: reqBody.Owner,
	}

	for i := range todos {
		todos[i].Id = uuid.New().String()
		dbList.Todos = append(dbList.Todos, todos[i])
	}

	err := db.lists.Create(&dbList).Error

	if err != nil {
		return nil, err
	}

	return &dbList, nil
}

func (db *ToDoListDB) GetList(listId string) (*models.ToDoList, error) {
	var list models.ToDoList

	result := db.lists.Preload("Todos").First(&list, "id = ?", listId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	return &list, nil
}

func (db *ToDoListDB) GetLists(username string) ([]*models.ToDoList, error) {
	var lists []*models.ToDoList

	result := db.lists.Preload("Todos").Find(&lists, "owner = ?", username)

	if result.Error != nil {
		return nil, result.Error
	}

	return lists, nil
}

func (db *ToDoListDB) GetTodos(listId string, username string) (*models.ToDoList, error) {
	var list models.ToDoList

	result := db.lists.Preload("Todos").First(&list, "id = ?", listId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, result.Error
	}

	return &list, nil
}

func (db *ToDoListDB) PatchList(reqBody *models.RequestBodyList, listId string, username string, role string) (*models.ToDoList, error) {

	var list models.ToDoList

	result := db.lists.Preload("Todos").First(&list, "id = ?", listId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	if list.Owner != username && role != "admin" {
		return nil, errors.New("action not allowed")
	}

	list.Owner = reqBody.Owner

	if err := db.lists.Table("to_do_lists").Where("id = ?", list.Id).Update("owner", list.Owner).Error; err != nil {
		return nil, err
	}

	return &list, nil
}

func (db *ToDoListDB) DeleteList(listId string, username string, role string) error {

	var list models.ToDoList

	result := db.lists.First(&list, "id = ?", listId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}

		return result.Error
	}

	if list.Owner != username && role != "admin" {
		return errors.New("action not allowed")
	}

	if err := db.lists.Delete(&list).Error; err != nil {
		return err
	}

	return nil
}

func (db *ToDoListDB) DeleteAllLists() error {
	err := db.lists.Where("1=1").Delete(&models.ToDoList{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (db *ToDoListDB) CreateTodo(reqBody *models.ToDo, listId string, username string) (*models.ToDo, error) {

	var list models.ToDoList

	result := db.lists.First(&list, "id = ?", listId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, result.Error
	}

	if list.Owner != username {
		return nil, errors.New("action not allowed")
	}

	todo := &models.ToDo{
		Id:        uuid.New().String(),
		ListId:    listId,
		Content:   reqBody.Content,
		Completed: false,
	}

	result = db.lists.Create(&todo)

	if result.Error != nil {
		return nil, result.Error
	}

	return todo, nil

}

func (db *ToDoListDB) GetTodo(todoId string, username string) (*models.ToDo, error) {
	var todo models.ToDo
	var list models.ToDoList

	result := db.lists.First(&todo, "id = ?", todoId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, result.Error
	}

	err := db.lists.First(&list, "id = ?", todo.ListId).Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	if list.Owner != username {
		return nil, errors.New("action not allowed")
	}

	return &todo, nil
}

func (db *ToDoListDB) PatchTodo(reqBody *models.ToDo, todoId string, username string) (*models.ToDo, error) {
	var todo models.ToDo
	var list models.ToDoList

	result := db.lists.First(&todo, "id = ?", todoId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, result.Error
	}

	err := db.lists.First(&list, "id = ?", todo.ListId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	if list.Owner != username {
		return nil, errors.New("action not allowed")
	}

	todo.Completed = reqBody.Completed

	if err := db.lists.Table("to_dos").Where("id = ?", todo.Id).Update("completed", todo.Completed).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (db *ToDoListDB) DeleteTodo(todoId string, username string) error {

	var todo models.ToDo
	var list models.ToDoList

	result := db.lists.First(&todo, "id = ?", todoId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}

		return result.Error
	}

	err := db.lists.First(&list, "id = ?", todo.ListId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}

		return err
	}

	if list.Owner != username {
		return errors.New("action not allowed")
	}

	todoErr := db.lists.Delete(&todo).Error

	if todoErr != nil {
		return todoErr
	}

	return nil
}

func (db *ToDoListDB) CreateUser(reqBody *models.User) (*models.User, error) {

	err := db.lists.Create(reqBody).Error

	if err != nil {
		return nil, err
	}

	return reqBody, nil
}

func (db *ToDoListDB) Login(reqBody *models.User) (*models.User, error) {

	var user models.User

	err := db.lists.First(&user, "username = ? AND password = ?", reqBody.Username, reqBody.Password).Error
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return &user, nil

}

func (db *ToDoListDB) DeleteUser(username string) error {

	result := db.lists.Where("username = ?", username).Delete(&models.User{})

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}

		return result.Error
	}

	return nil

}

func (db *ToDoListDB) GetAllListsAdmin() ([]*models.ToDoList, error) {
	var lists []*models.ToDoList

	result := db.lists.Preload("Todos").Find(&lists)

	if result.Error != nil {
		return nil, result.Error
	}

	return lists, nil
}
