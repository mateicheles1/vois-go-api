package controllers

import (
	"gogin-api/models"
	"gogin-api/service"

	"github.com/gin-gonic/gin"
)

func Lists(c *gin.Context) {

	lists, err := service.Repo.GetAllLists()
	if err != nil {
		c.Status(204)
		return
	}
	c.JSON(200, lists)

}

func Todos(c *gin.Context) {

	todos, err := service.Repo.GetAllToDosInList(c.Param("listid"))

	if err != nil {
		c.Status(404)
		return
	}

	c.JSON(200, todos)

}

func GetList(c *gin.Context) {

	list, err := service.Repo.GetList(c.Param("listid"))

	if err != nil {
		c.Status(404)
		return
	}

	c.JSON(200, list)

}

func CreateList(c *gin.Context) {
	requestBody := new(models.RequestBodyList)

	if err := c.BindJSON(requestBody); err != nil {
		return
	}

	service.Repo.CreateList(requestBody)

	c.Status(201)

}

func PatchList(c *gin.Context) {

	requestBody := new(models.ToDoList)

	if err := c.BindJSON(requestBody); err != nil {
		return
	}

	err := service.Repo.PatchList(requestBody)

	if err != nil {
		c.Status(404)
		return
	}

	c.Status(200)

}

func DeleteList(c *gin.Context) {

	err := service.Repo.DeleteList(c.Param("listid"))

	if err != nil {
		c.Status(404)
		return
	}

	c.Status(200)

}

func GetToDo(c *gin.Context) {

	todo, err := service.Repo.GetToDoInList(c.Param("todoid"))

	if err != nil {
		c.Status(404)
		return
	}

	c.JSON(200, todo)

}

func DeleteToDo(c *gin.Context) {

	err := service.Repo.DeleteToDoInList(c.Param("todoid"))

	if err != nil {
		c.Status(404)
		return
	}

	c.Status(200)

}

func PatchToDo(c *gin.Context) {
	requestBody := new(models.ToDo)

	if err := c.BindJSON(requestBody); err != nil {
		return
	}

	err := service.Repo.PatchToDoInList(requestBody.Content, c.Param("todoid"))

	if err != nil {
		c.Status(404)
		return
	}

	c.Status(200)
}

func CreateToDo(c *gin.Context) {

	requestBody := new(models.ToDo)

	if err := c.BindJSON(requestBody); err != nil {
		return
	}

	err := service.Repo.CreateToDoInList(c.Param("listid"), requestBody.Content)

	if err != nil {
		c.Status(404)
		return
	}

	c.Status(201)
}
