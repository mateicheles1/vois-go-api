package controllers

import (
	"gogin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllListsHandler(todoListService models.ToDoListService) gin.HandlerFunc {
	return func(c *gin.Context) {
		lists, err := todoListService.GetAllLists()
		if err != nil {
			c.Status(204)
			return
		}
		c.JSON(http.StatusOK, lists)
	}

}

func GetAllTodosHandler(todoListService models.ToDoListService) gin.HandlerFunc {
	return func(c *gin.Context) {

		todos, err := todoListService.GetAllToDosInList(c.Param("listid"))

		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, todos)
	}

}

func GetListHandler(todoListService models.ToDoListService) gin.HandlerFunc {

	return func(c *gin.Context) {

		list, err := todoListService.GetList(c.Param("listid"))

		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, list)
	}

}

func CreateListHandler(todoListService models.ToDoListService) gin.HandlerFunc {

	return func(c *gin.Context) {

		requestBody := new(models.RequestBodyList)

		if err := c.BindJSON(requestBody); err != nil {
			return
		}

		todoListService.CreateList(requestBody)

		c.Status(http.StatusCreated)
	}

}

func PatchListHandler(todoListService models.ToDoListService) gin.HandlerFunc {

	return func(c *gin.Context) {

		requestBody := new(models.ToDoList)

		if err := c.BindJSON(requestBody); err != nil {
			return
		}

		requestBody.Id = c.Param("listid")

		err := todoListService.PatchList(requestBody)

		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusOK)
	}

}

func DeleteListHandler(todoListService models.ToDoListService) gin.HandlerFunc {

	return func(c *gin.Context) {

		err := todoListService.DeleteList(c.Param("listid"))

		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusOK)
	}

}

func GetToDoHandler(todoListService models.ToDoListService) gin.HandlerFunc {

	return func(c *gin.Context) {

		todo, err := todoListService.GetToDoInList(c.Param("todoid"))

		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, todo)
	}

}

func DeleteToDoHandler(todoListService models.ToDoListService) gin.HandlerFunc {

	return func(c *gin.Context) {

		err := todoListService.DeleteToDoInList(c.Param("todoid"))

		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusOK)
	}

}

func PatchToDoHandler(todoListService models.ToDoListService) gin.HandlerFunc {

	return func(c *gin.Context) {

		requestBody := new(models.ToDo)

		if err := c.BindJSON(requestBody); err != nil {
			return
		}

		err := todoListService.PatchToDoInList(requestBody.Completed, c.Param("todoid"))

		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusOK)
	}
}

func CreateToDoHandler(todoListService models.ToDoListService) gin.HandlerFunc {

	return func(c *gin.Context) {

		requestBody := new(models.ToDo)

		if err := c.BindJSON(requestBody); err != nil {
			return
		}
		if requestBody.Content == "" {
			c.String(http.StatusBadRequest, "content can't be empty")
			return
		}
		err := todoListService.CreateToDoInList(c.Param("listid"), requestBody.Content)

		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusCreated)
	}

}
