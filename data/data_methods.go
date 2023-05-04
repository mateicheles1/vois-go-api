package data

import (
	"errors"
	"gogin-api/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ToDoListDB struct {
	Lists *gorm.DB
}

func NewToDoListDB(db *gorm.DB) *ToDoListDB {
	return &ToDoListDB{
		Lists: db,
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

	err := db.Lists.Create(&dbList).Error

	if err != nil {
		return nil, err
	}

	resBodyList := &models.ToDoList{
		Id:    dbList.Id,
		Owner: dbList.Owner,
		Todos: make([]*models.ToDo, len(dbList.Todos)),
	}

	for i, todo := range dbList.Todos {
		resBodyList.Todos[i] = &models.ToDo{
			Id:        todo.Id,
			Content:   todo.Content,
			Completed: todo.Completed,
		}
	}

	return resBodyList, nil
}

func (db *ToDoListDB) GetList(listId string) (*models.ToDoList, error) {
	var list models.ToDoList

	result := db.Lists.Preload("Todos").First(&list, "id = ?", listId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	return &list, nil
}

func (db *ToDoListDB) GetLists() ([]*models.ToDoList, error) {
	var lists []*models.ToDoList

	result := db.Lists.Preload("Todos").Find(&lists)

	if result.Error != nil {
		return nil, result.Error
	}

	return lists, nil
}

func (db *ToDoListDB) GetTodos(listId string) ([]*models.ToDo, error) {
	var todos []*models.ToDo

	result := db.Lists.Find(&todos, "list_id = ?", listId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, result.Error
	}

	return todos, nil
}

func (db *ToDoListDB) PatchList(reqBody *models.RequestBodyList, listId string) (*models.ToDoList, error) {

	var list models.ToDoList

	result := db.Lists.Preload("Todos").First(&list, "id = ?", listId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	list.Owner = reqBody.Owner

	if err := db.Lists.Model(&list).Update("owner", list.Owner).Error; err != nil {
		return nil, err
	}

	return &list, nil
}

func (db *ToDoListDB) DeleteList(listId string) error {

	var list models.ToDoList

	result := db.Lists.First(&list, "id = ?", listId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}

		return result.Error
	}

	if err := db.Lists.Delete(&list).Error; err != nil {
		return err
	}

	return nil
}

func (db *ToDoListDB) CreateTodo(reqBody *models.ToDo, listId string) (*models.ToDo, error) {

	result := db.Lists.First(&models.ToDoList{}, "id = ?", listId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, result.Error
	}

	todo := &models.ToDo{
		Id:        uuid.New().String(),
		ListId:    listId,
		Content:   reqBody.Content,
		Completed: false,
	}

	result = db.Lists.Create(todo)

	if result.Error != nil {
		return nil, result.Error
	}

	return todo, nil

}

func (db *ToDoListDB) GetTodo(todoId string) (*models.ToDo, error) {
	var todo models.ToDo

	result := db.Lists.First(&todo, "id = ?", todoId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, result.Error
	}

	return &todo, nil
}

func (db *ToDoListDB) PatchTodo(reqBody *models.ToDo, todoId string) (*models.ToDo, error) {
	var todo models.ToDo

	result := db.Lists.First(&todo, "id = ?", todoId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, result.Error
	}

	todo.Completed = reqBody.Completed

	if err := db.Lists.Model(&todo).Update("completed", todo.Completed).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (db *ToDoListDB) DeleteTodo(todoId string) error {
	var todo models.ToDo

	result := db.Lists.First(&todo, "id = ?", todoId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}

		return result.Error
	}

	if err := db.Lists.Delete(&todo).Error; err != nil {
		return err
	}

	return nil
}
