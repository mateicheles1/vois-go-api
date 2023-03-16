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
	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")].Todos)
}

func getList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")])
}

func createList(c *gin.Context) {
	requestBody := new(models.ToDoList)
	mapCopyRequestBody := make(map[string]*models.ToDoList)
	bodyTodos := make(map[string]*models.ToDo)
	requestBodyKey := uuid.New().String()
	
	
	if err := c.BindJSON(requestBody); err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	if requestBody.Owner == "" || requestBody.Todos == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid content"})
	}
	
	for k := range requestBody.Todos {
		toDosKey := uuid.New().String()
		requestBody.Todos[k].Id = toDosKey
		requestBody.Todos[k].Listid = requestBodyKey
		bodyTodos[toDosKey] = requestBody.Todos[k]
	}
	requestBody.Id = requestBodyKey
	mapCopyRequestBody[requestBodyKey] = requestBody
	mapCopyRequestBody[requestBodyKey].Todos = bodyTodos

	models.Data[requestBodyKey] = mapCopyRequestBody[requestBodyKey]
	c.IndentedJSON(http.StatusOK, models.Data)

}


func updateList(c *gin.Context) {
	requestBody := new(models.ToDoList)

	if err := c.BindJSON(requestBody); err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	if requestBody.Owner == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid owner"})
	}
	models.Data[c.Param("listid")].Owner = requestBody.Owner
	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")])
}

func deleteList(c *gin.Context) {
	delete(models.Data, c.Param("listid"))
	c.IndentedJSON(http.StatusOK, models.Data)
}


func getToDo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK,models.Data[c.Param("listid")].Todos[c.Param("todoid")])
}

func deleteToDo(c *gin.Context) {
	delete(models.Data[c.Param("listid")].Todos, c.Param("todoid"))
	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")].Todos)
}

func updateToDo(c *gin.Context) {
	requestBody := new(models.ToDo)

	if err := c.BindJSON(requestBody); err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	if requestBody.Content == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid content"})
	}


	models.Data[c.Param("listid")].Todos[c.Param("todoid")].Content = requestBody.Content
	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")].Todos[c.Param("todoid")])
}

func createToDo(c *gin.Context) {
	requestBody := new(models.ToDo)

	if err := c.BindJSON(requestBody); err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	if requestBody.Content == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid content"})
	}

	key := uuid.New().String()
	requestBody.Listid = c.Param("listid")
	requestBody.Id = key
	models.Data[c.Param("listid")].Todos[key] = requestBody
	c.IndentedJSON(http.StatusOK, models.Data[c.Param("listid")])
}