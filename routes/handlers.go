package routes

import (
	"gogin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func lists(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Data)
}

func todos(c *gin.Context) {
	_, hasList := models.Data[c.Param("listid")]


		if !hasList {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "list not found"})
		} else {
			c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")].Todos)
		}
}

func getList(c *gin.Context) {
	_, hasList := models.Data[c.Param("listid")]


		if !hasList {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "list not found"})
		} else {
			c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")])
		}

}

func createList(c *gin.Context) {
	requestBody := new(models.RequestBodyList)
	requestBodyTodos := make(map[string]*models.ToDo)
	todoListKey := uuid.New().String()

	if err := c.ShouldBindJSON(requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		check(err, c)
	}

	if len(requestBody.Todos) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "todos can't be empty"})
		return
	}

	for _, v := range requestBody.Todos {
		toDosKey := uuid.New().String()
		toDo := new(models.ToDo)
		toDo.Content = v
		toDo.Id = toDosKey
		requestBodyTodos[toDosKey] = toDo
	}

	models.Data[todoListKey] = new(models.ToDoList)
	models.Data[todoListKey].Id = todoListKey
	models.Data[todoListKey].Owner = requestBody.Owner
	models.Data[todoListKey].Todos = requestBodyTodos

	c.IndentedJSON(http.StatusOK, models.Data[todoListKey])
}

func updateList(c *gin.Context) {

	_, hasList := models.Data[c.Param("listid")]
	if !hasList {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "list not found"})
		return
	}

	requestBody := new(models.ToDoList)

	if err := c.ShouldBindJSON(requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		check(err, c)
	}

	models.Data[c.Param("listid")].Owner = requestBody.Owner

	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")])
}


func deleteList(c *gin.Context) {

	_, hasList := models.Data[c.Param("listid")]


		if !hasList {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "list not found"})
		} else {
			delete(models.Data, c.Param("listid"))
		}


}


func getToDo(c *gin.Context) {
	_, hasList := models.Data[c.Param("listid")]


		if !hasList {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "list not found"})
		}

		if hasList {
			_, hasTodo := models.Data[c.Param("listid")].Todos[c.Param("todoid")]
			if !hasTodo {
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
			} else {
				c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")].Todos[c.Param("todoid")])
			}
		}

}


func deleteToDo(c *gin.Context) {
	
	_, hasList := models.Data[c.Param("listid")]


		if !hasList {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "list not found"})
		}

		if hasList {
			_, hasTodo := models.Data[c.Param("listid")].Todos[c.Param("todoid")]
			if !hasTodo {
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
			} else {
				delete(models.Data[c.Param("listid")].Todos, c.Param("todoid"))
			}
		}

}

func updateToDo(c *gin.Context) {
	_, hasList := models.Data[c.Param("listid")]

		if !hasList {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "list not found"})
			return
		} else {
			_, hasTodo := models.Data[c.Param("listid")].Todos[c.Param("todoid")]
			if !hasTodo {
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
				return
			}
		}

	requestBody := new(models.ToDo)

		if err := c.ShouldBindJSON(requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		check(err, c)
	}


	models.Data[c.Param("listid")].Todos[c.Param("todoid")].Content = requestBody.Content


	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")].Todos[c.Param("todoid")])
}

func createToDo(c *gin.Context) {
	_, hasList := models.Data[c.Param("listid")]


		if !hasList {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "list not found"})
			return
		}

	requestBody := new(models.ToDo)

	if err := c.ShouldBindJSON(requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		check(err, c)
	}

	key := uuid.New().String()
	requestBody.Id = key
	models.Data[c.Param("listid")].Todos[key] = requestBody


	c.IndentedJSON(http.StatusCreated, models.Data[c.Param("listid")].Todos[key])
}