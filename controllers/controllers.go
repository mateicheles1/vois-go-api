package controllers

import (
	"gogin-api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Lists(c *gin.Context) {
	if len(models.Data) == 0 {
		c.Status(204)
	} else {
		var lists []*models.ToDoList
		for _, list := range models.Data {
			lists = append(lists, list)
		}
		c.JSON(200, lists)
	}
}

func Todos(c *gin.Context) {
	list, hasList := models.Data[c.Param("listid")]

	if !hasList {
		c.Status(404)
	} else {
		var todos []*models.ToDo
		for _, todo := range list.Todos {
			todos = append(todos, todo)
		}
		c.JSON(200, todos)
	}

}

func GetList(c *gin.Context) {
	_, hasList := models.Data[c.Param("listid")]

	if !hasList {
		c.Status(404)
	} else {
		c.JSON(200, models.Data[c.Param("listid")])
	}

}

func CreateList(c *gin.Context) {
	requestBody := new(models.RequestBodyList)
	requestBodyTodos := make(map[string]*models.ToDo)
	todoListKey := uuid.New().String()

	if err := c.BindJSON(requestBody); err != nil {
		c.AbortWithError(400, err)
		return
	}

	for _, v := range requestBody.Todos {
		toDosKey := uuid.New().String()
		requestBodyTodos[toDosKey] = &models.ToDo{
			Id:      toDosKey,
			Content: v,
		}
	}

	models.Data[todoListKey] = &models.ToDoList{
		Id:    todoListKey,
		Owner: requestBody.Owner,
		Todos: requestBodyTodos,
	}

	c.Status(200)
}

func PatchList(c *gin.Context) {

	_, hasList := models.Data[c.Param("listid")]
	if !hasList {
		c.Status(404)
		return
	}

	requestBody := new(models.ToDoList)

	if err := c.BindJSON(requestBody); err != nil {
		c.AbortWithError(400, err)
		return
	}

	models.Data[c.Param("listid")].Owner = requestBody.Owner

	c.Status(200)
}

func DeleteList(c *gin.Context) {

	_, hasList := models.Data[c.Param("listid")]

	if !hasList {
		c.Status(404)
	} else {
		delete(models.Data, c.Param("listid"))
	}

}

func GetToDo(c *gin.Context) {
	for _, list := range models.Data {
		todo, hasToDo := list.Todos[c.Param("todoid")]

		if !hasToDo {
			c.Status(404)
		} else {
			c.JSON(200, todo)
			break
		}
	}
}

func DeleteToDo(c *gin.Context) {

	for k, list := range models.Data {
		_, hasToDo := list.Todos[c.Param("todoid")]

		if !hasToDo {
			c.Status(404)
		} else {
			delete(models.Data[k].Todos, c.Param("todoid"))
			break
		}
	}

}

func PatchToDo(c *gin.Context) {
	requestBody := new(models.ToDo)

	if err := c.BindJSON(requestBody); err != nil {
		c.AbortWithError(400, err)
		return
	}

	for k := range models.Data {
		_, hasToDo := models.Data[k].Todos[c.Param("todoid")]

		if !hasToDo {
			c.Status(404)
		} else {
			models.Data[k].Todos[c.Param("todoid")].Content = requestBody.Content
			c.Status(200)
			break
		}
	}
}

func CreateToDo(c *gin.Context) {
	_, hasList := models.Data[c.Param("listid")]

	if !hasList {
		c.Status(404)
		return
	}

	requestBody := new(models.ToDo)

	if err := c.BindJSON(requestBody); err != nil {
		c.AbortWithError(400, err)
		return
	}

	key := uuid.New().String()

	models.Data[c.Param("listid")].Todos[key] = &models.ToDo{
		Id:      key,
		Content: requestBody.Content,
	}

	c.Status(200)
}
