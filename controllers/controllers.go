package controllers

import (
	"gogin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Lists(c *gin.Context) {
	if len(models.Data) == 0 {
		c.Status(http.StatusNoContent)
	} else {
		c.IndentedJSON(http.StatusOK, models.Data)
	}
}

func Todos(c *gin.Context) {
	_, hasList := models.Data[c.Param("listid")]

		if !hasList {
			c.Status(http.StatusNotFound)
		} else {
			var todos []*models.ToDo
			for _, todo := range models.Data[c.Param("listid")].Todos {
				todos = append(todos, todo)
			}
			c.IndentedJSON(http.StatusOK, todos)
		}

}

func GetList(c *gin.Context) {
	_, hasList := models.Data[c.Param("listid")]


		if !hasList {
			c.Status(http.StatusNotFound)
		} else {
			c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")])
		}

}

func CreateList(c *gin.Context) {
	requestBody := new(models.RequestBodyList)
	requestBodyTodos := make(map[string]*models.ToDo)
	todoListKey := uuid.New().String()

	if err := c.ShouldBindJSON(requestBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for _, v := range requestBody.Todos {
		toDosKey := uuid.New().String()
		requestBodyTodos[toDosKey] = &models.ToDo{
			Content: v,
			Id: toDosKey,
		}
	}

	models.Data[todoListKey] = &models.ToDoList{
		Id: todoListKey,
		Owner: requestBody.Owner,
		Todos: requestBodyTodos,
	}

	c.IndentedJSON(http.StatusCreated, models.Data[todoListKey])
}

func PatchList(c *gin.Context) {

	_, hasList := models.Data[c.Param("listid")]
	if !hasList {
		c.Status(http.StatusNotFound)
		return
	}

	requestBody := new(models.ToDoList)

	if err := c.ShouldBindJSON(requestBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	models.Data[c.Param("listid")].Owner = requestBody.Owner

	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")])
}


func DeleteList(c *gin.Context) {

	_, hasList := models.Data[c.Param("listid")]


		if !hasList {
			c.Status(http.StatusNotFound)
		} else {
			delete(models.Data, c.Param("listid"))
		}


}


func GetToDo(c *gin.Context) {
	_, hasList := models.Data[c.Param("listid")]


		if !hasList {
			c.IndentedJSON(http.StatusNotFound, "404 list not found")
		} else {
			_, hasTodo := models.Data[c.Param("listid")].Todos[c.Param("todoid")]
			if !hasTodo {
				c.IndentedJSON(http.StatusNotFound, "404 todo not found")
			} else {
				c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")].Todos[c.Param("todoid")])
			}
		}

}


func DeleteToDo(c *gin.Context) {
	
	_, hasList := models.Data[c.Param("listid")]


		if !hasList {
			c.IndentedJSON(http.StatusNotFound, "404 list not found")
		} else {
			_, hasTodo := models.Data[c.Param("listid")].Todos[c.Param("todoid")]
			if !hasTodo {
				c.IndentedJSON(http.StatusNotFound, "404 todo not found")
			} else {
				delete(models.Data[c.Param("listid")].Todos, c.Param("todoid"))
			}
		}

}

func PatchToDo(c *gin.Context) {
	_, hasList := models.Data[c.Param("listid")]

		if !hasList {
			c.IndentedJSON(http.StatusNotFound, "404 list not found")
			return
		} else {
			_, hasTodo := models.Data[c.Param("listid")].Todos[c.Param("todoid")]
			if !hasTodo {
				c.IndentedJSON(http.StatusNotFound, "404 todo not found")
				return
			}
		}

	requestBody := new(models.ToDo)

		if err := c.ShouldBindJSON(requestBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}


	models.Data[c.Param("listid")].Todos[c.Param("todoid")].Content = requestBody.Content


	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")].Todos[c.Param("todoid")])
}

func CreateToDo(c *gin.Context) {
	_, hasList := models.Data[c.Param("listid")]


		if !hasList {
			c.IndentedJSON(http.StatusNotFound, "404 list not found")
			return
		}

	requestBody := new(models.ToDo)

	if err := c.ShouldBindJSON(requestBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	key := uuid.New().String()
	
	models.Data[c.Param("listid")].Todos[key] = &models.ToDo{
		Id: key,
		Content: requestBody.Content,
	}


	c.IndentedJSON(http.StatusCreated, models.Data[c.Param("listid")].Todos[key])
}