package controllers

import (
	"gogin-api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Lists(c *gin.Context) {

	if models.AllData.Lists == nil {
		c.Status(204)
	} else {
		c.JSON(200, models.AllData.PrintAllLists())
	}

}

func Todos(c *gin.Context) {

	if list, hasList := models.AllData.Lists[c.Param("listid")]; !hasList {
		c.Status(404)
	} else {
		c.JSON(200, list.PrintTodos())
	}

}

func GetList(c *gin.Context) {

	if list, hasList := models.AllData.Lists[c.Param("listid")]; !hasList {
		c.Status(404)
	} else {
		c.JSON(200, list.PrintList())
	}

}

func CreateList(c *gin.Context) {
	requestBody := new(models.RequestBodyList)
	requestBodyTodos := make(map[string]*models.ToDo)
	todoListKey := uuid.New().String()

	if err := c.BindJSON(requestBody); err != nil {
		return
	}

	for _, v := range requestBody.Todos {
		toDosKey := uuid.New().String()
		requestBodyTodos[toDosKey] = &models.ToDo{
			Id:      toDosKey,
			ListId:  todoListKey,
			Content: v,
		}
	}

	models.AllData.CreateList(requestBody, requestBodyTodos, todoListKey)
	c.Status(201)

}

func PatchList(c *gin.Context) {

	if _, hasList := models.AllData.Lists[c.Param("listid")]; !hasList {
		c.Status(404)
		return
	}

	requestBody := new(models.ToDoList)

	if err := c.BindJSON(requestBody); err != nil {
		return
	}

	models.AllData.Lists[c.Param("listid")].PatchList(requestBody.Owner)
	c.Status(200)
}

func DeleteList(c *gin.Context) {

	if _, hasList := models.AllData.Lists[c.Param("listid")]; !hasList {
		c.Status(404)
	} else {
		models.AllData.DeleteList(c.Param("listid"))
	}

}

func GetToDo(c *gin.Context) {

	for _, list := range models.AllData.Lists {
		if todo, hasToDo := list.Todos[c.Param("todoid")]; !hasToDo {
			c.Status(404)
		} else {
			c.JSON(200, todo.PrintToDo())
			break
		}
	}

}

func DeleteToDo(c *gin.Context) {

	for k, list := range models.AllData.Lists {
		if _, hasToDo := list.Todos[c.Param("todoid")]; !hasToDo {
			c.Status(404)
		} else {
			models.AllData.Lists[k].DeleteToDo(c.Param("todoid"))
			break
		}
	}

}

func PatchToDo(c *gin.Context) {
	requestBody := new(models.ToDo)

	if err := c.BindJSON(requestBody); err != nil {
		return
	}

	for k, list := range models.AllData.Lists {
		if _, hasToDo := list.Todos[c.Param("todoid")]; !hasToDo {
			c.Status(404)
			return
		} else {
			models.AllData.Lists[k].Todos[c.Param("todoid")].PatchToDo(requestBody.Content)
			c.Status(200)
			break
		}
	}
}

func CreateToDo(c *gin.Context) {

	if _, hasList := models.AllData.Lists[c.Param("listid")]; !hasList {
		c.Status(404)
		return
	}

	requestBody := new(models.ToDo)

	if err := c.BindJSON(requestBody); err != nil {
		return
	}

	models.AllData.Lists[c.Param("listid")].CreateToDo(requestBody)
	c.Status(201)
}
