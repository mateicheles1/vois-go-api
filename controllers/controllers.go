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
		c.JSON(200, models.Data)
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
			Id: toDosKey,
			Content: v,
		}
	}

	models.Data[todoListKey] = &models.ToDoList{
		Id: todoListKey,
		Owner: requestBody.Owner,
		Todos: requestBodyTodos,
	}

	c.JSON(201, models.Data[todoListKey])
}

func PatchList(c *gin.Context) {

	_, hasList := models.Data[c.Param("listid")]
	if !hasList {
		c.Status(404)
		return
	}

	requestBody := new(models.ToDoList)

	if err := c.ShouldBindJSON(requestBody); err != nil {
		c.AbortWithError(400, err)
		return
	}

	models.Data[c.Param("listid")].Owner = requestBody.Owner

	c.JSON(200, models.Data[c.Param("listid")])
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
	list, hasList := models.Data[c.Param("listid")]

		if !hasList {
			c.JSON(404, "404 list not found")
		} else {
			todo, hasToDo := list.Todos[c.Param("todoid")]
			if !hasToDo {
				c.JSON(404, "404 todo not found")
			} else {
				c.JSON(200, todo)
			}
	}
}



func DeleteToDo(c *gin.Context) {
	
	_, hasList := models.Data[c.Param("listid")]


		if !hasList {
			c.JSON(404, "404 list not found")
		} else {
			_, hasTodo := models.Data[c.Param("listid")].Todos[c.Param("todoid")]
			if !hasTodo {
				c.JSON(404, "404 todo not found")
			} else {
				delete(models.Data[c.Param("listid")].Todos, c.Param("todoid"))
			}
		}

}

func PatchToDo(c *gin.Context) {
	_, hasList := models.Data[c.Param("listid")]

		if !hasList {
			c.JSON(404, "404 list not found")
			return
		} else {
			_, hasTodo := models.Data[c.Param("listid")].Todos[c.Param("todoid")]
			if !hasTodo {
				c.JSON(404, "404 todo not found")
				return
			}
		}

	requestBody := new(models.ToDo)

		if err := c.ShouldBindJSON(requestBody); err != nil {
		c.AbortWithError(400, err)
		return
	}


	models.Data[c.Param("listid")].Todos[c.Param("todoid")].Content = requestBody.Content


	c.JSON(200, models.Data[c.Param("listid")].Todos[c.Param("todoid")])
}

func CreateToDo(c *gin.Context) {
	_, hasList := models.Data[c.Param("listid")]


		if !hasList {
			c.JSON(404, "404 list not found")
			return
		}

	requestBody := new(models.ToDo)

	if err := c.ShouldBindJSON(requestBody); err != nil {
		c.AbortWithError(400, err)
		return
	}

	key := uuid.New().String()
	
	models.Data[c.Param("listid")].Todos[key] = &models.ToDo{
		Id: key,
		Content: requestBody.Content,
	}


	c.JSON(201, models.Data[c.Param("listid")].Todos[key])
}