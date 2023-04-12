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
	
	if list, hasList := models.Data[c.Param("listid")]; !hasList {
		c.Status(404)
	} else {
			var todos []*models.ToDo
			for _, todo := range list.Todos {
				responseToDo := &models.ToDo{
					Id: todo.Id,
					Content: todo.Content,
				}
				todos = append(todos, responseToDo)
			}
			c.JSON(200, todos)
		}

}

func GetList(c *gin.Context) {

	if list, hasList := models.Data[c.Param("listid")]; !hasList {
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
			Id: toDosKey,
			ListId: todoListKey,
			Content: v,
		}
	}

	models.Data[todoListKey] = &models.ToDoList{
		Id: todoListKey,
		Owner: requestBody.Owner,
		Todos: requestBodyTodos,
	}

	c.Status(201)
}

func PatchList(c *gin.Context) {

	if _, hasList := models.Data[c.Param("listid")]; !hasList {
		c.Status(404)
		return
	}

	requestBody := new(models.ToDoList)

	if err := c.BindJSON(requestBody); err != nil {
		return
	}

	models.Data[c.Param("listid")].Owner = requestBody.Owner

	c.Status(200)
}


func DeleteList(c *gin.Context) {

	if _, hasList := models.Data[c.Param("listid")]; !hasList {
		c.Status(404)
	} else {
			delete(models.Data, c.Param("listid"))
		}


}


func GetToDo(c *gin.Context) {

	for _, list := range models.Data {
		if todo, hasToDo := list.Todos[c.Param("todoid")]; !hasToDo {
			c.Status(404)
		} else {
			c.JSON(200, todo.PrintToDo())
			break
		}
	}

}



func DeleteToDo(c *gin.Context) {
	
	for k, list := range models.Data {
		if _, hasToDo := list.Todos[c.Param("todoid")]; !hasToDo {
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
		return
	}

	for k, list := range models.Data {
		if _, hasToDo := list.Todos[c.Param("todoid")]; !hasToDo {
			c.Status(404)
			return
		} else {
			models.Data[k].Todos[c.Param("todoid")].Content = requestBody.Content
			c.Status(200)
			break
		}
	}

}

func CreateToDo(c *gin.Context) {

	if _, hasList := models.Data[c.Param("listid")]; !hasList {
		c.Status(404)
		return
	}

	requestBody := new(models.ToDo)

	if err := c.BindJSON(requestBody); err != nil {
		return
	}

	key := uuid.New().String()
	
	models.Data[c.Param("listid")].Todos[key] = &models.ToDo{
		Id: key,
		ListId: c.Param("listid"),
		Content: requestBody.Content,
	}

	c.Status(201)
}