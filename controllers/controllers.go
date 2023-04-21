package controllers

import (
	"gogin-api/models"

	"github.com/gin-gonic/gin"
)

func GetAllListsHandler(todoListService models.ToDoListService) gin.HandlerFunc {
	return func(c *gin.Context) {
		lists, err := todoListService.GetAllLists()
		if err != nil {
			c.Status(204)
			return
		}
		c.JSON(200, lists)
	}

}

func GetAllTodosHandler(todoListService models.ToDoListService) gin.HandlerFunc {
	return func(c *gin.Context) {

		todos, err := todoListService.GetAllToDosInList(c.Param("listid"))

		if err != nil {
			c.Status(404)
			return
		}

		c.JSON(200, todos)
	}

}

func GetListHandler(todoListService models.ToDoListService) gin.HandlerFunc {

	return func(c *gin.Context) {

		list, err := todoListService.GetList(c.Param("listid"))

		if err != nil {
			c.Status(404)
			return
		}

		c.JSON(200, list)
	}

}

func CreateListHandler(todoListService models.ToDoListService) gin.HandlerFunc {

	return func(c *gin.Context) {

		requestBody := new(models.RequestBodyList)

		if err := c.BindJSON(requestBody); err != nil {
			return
		}

		todoListService.CreateList(requestBody)

		c.Status(201)
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
			c.Status(404)
			return
		}

		c.Status(200)
	}

}

func DeleteListHandler(todoListService models.ToDoListService) gin.HandlerFunc {

	return func(c *gin.Context) {

		err := todoListService.DeleteList(c.Param("listid"))

		if err != nil {
			c.Status(404)
			return
		}

		c.Status(200)
	}

}

func GetToDoHandler(todoListService models.ToDoListService) gin.HandlerFunc {

	return func(c *gin.Context) {

		todo, err := todoListService.GetToDoInList(c.Param("todoid"))

		if err != nil {
			c.Status(404)
			return
		}

		c.JSON(200, todo)
	}

}

func DeleteToDoHandler(todoListService models.ToDoListService) gin.HandlerFunc {

	return func(c *gin.Context) {

		err := todoListService.DeleteToDoInList(c.Param("todoid"))

		if err != nil {
			c.Status(404)
			return
		}

		c.Status(200)
	}

}

func PatchToDoHandler(todoListService models.ToDoListService) gin.HandlerFunc {

	return func(c *gin.Context) {

		requestBody := new(models.ToDo)

		if err := c.BindJSON(requestBody); err != nil {
			return
		}

		err := todoListService.PatchToDoInList(requestBody.Content, c.Param("todoid"))

		if err != nil {
			c.Status(404)
			return
		}

		c.Status(200)
	}
}

func CreateToDoHandler(todoListService models.ToDoListService) gin.HandlerFunc {

	return func(c *gin.Context) {

		requestBody := new(models.ToDo)

		if err := c.BindJSON(requestBody); err != nil {
			return
		}

		err := todoListService.CreateToDoInList(c.Param("listid"), requestBody.Content)

		if err != nil {
			c.Status(404)
			return
		}

		c.Status(201)
	}

}
