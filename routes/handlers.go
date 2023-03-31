package routes

import (
	"gogin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func lists(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"data": models.Data})
}

func todos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"data": models.Data[c.Param("listid")].Todos})
}

func getList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"data": models.Data[c.Param("listid")]})
}

func createList(c *gin.Context) {
	// request body 
	requestBody := new(models.ToDoList)
	// un map de todo-uri care va primi pe key todo struct-ul din todolist struct-ul requestBody
	requestBodyTodos := make(map[string]*models.ToDo)
	// key din models.Data si id-ul struct-ului de todolist
	toDoListKey := uuid.New().String()
	
	
	if err := c.ShouldBindJSON(requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		check(err, c)
	}
	
	if requestBody.Todos == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "todos can't be empty"})
		return
	}
	

	for k := range requestBody.Todos {
		toDosKey := uuid.New().String()
		requestBody.Todos[k].Id = toDosKey
		requestBody.Todos[k].Listid = toDoListKey
		// map-ul map[toDosKey]*ToDo primeste struct-urile de todo-uri din request body
		requestBodyTodos[toDosKey] = requestBody.Todos[k]
	}
	
	requestBody.Id = toDoListKey
	models.Data[toDoListKey] = requestBody
	models.Data[toDoListKey].Todos = requestBodyTodos
	
 
	c.IndentedJSON(http.StatusOK, gin.H{"message": "list successfully created", "data": models.Data[toDoListKey]})

}


func updateList(c *gin.Context) {
	requestBody := new(models.ToDoList)

	if err := c.ShouldBindJSON(requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		check(err, c)
	}

	models.Data[c.Param("listid")].Owner = requestBody.Owner

	c.IndentedJSON(http.StatusOK, gin.H{"message": "list successfully updated", "data": models.Data[c.Param("listid")]})
}


func deleteList(c *gin.Context) {
	delete(models.Data, c.Param("listid"))


	c.IndentedJSON(http.StatusOK, gin.H{"message": "list successfully deleted"})
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
	requestBody := new(models.ToDo)

		if err := c.ShouldBindJSON(requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		check(err, c)
	}


	models.Data[c.Param("listid")].Todos[c.Param("todoid")].Content = requestBody.Content


	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")].Todos[c.Param("todoid")])
}

func createToDo(c *gin.Context) {
	requestBody := new(models.ToDo)

	if err := c.ShouldBindJSON(requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		check(err, c)
	}

	key := uuid.New().String()
	requestBody.Listid = c.Param("listid")
	requestBody.Id = key
	models.Data[c.Param("listid")].Todos[key] = requestBody


	c.IndentedJSON(http.StatusCreated, models.Data[c.Param("listid")].Todos[key])
}