package controllers

import (
	"fmt"
	"gogin-api/models"
	"gogin-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service service.ToDoListServiceInterface
}

func NewController(service service.ToDoListServiceInterface) *Controller {
	return &Controller{
		Service: service,
	}
}

func (c *Controller) CreateList(ctx *gin.Context) {
	var requestBody models.RequestBodyList

	if err := ctx.BindJSON(&requestBody); err != nil {
		return
	}

	list, err := c.Service.CreateList(requestBody)

	if err != nil {

		if err.Error() == "empty owner" {
			ctx.JSON(http.StatusBadRequest, "invalid JSON syntax in request body; empty owner")
			return
		}

		if err.Error() != "empty owner" {
			ctx.JSON(http.StatusInternalServerError, "something went wrong")
			return
		}
	}

	ctx.JSON(http.StatusCreated, list)
}

func (c *Controller) CreateToDo(ctx *gin.Context) {
	var requestBody models.ToDo

	if err := ctx.BindJSON(&requestBody); err != nil {
		return
	}

	listId := ctx.Param("listid")

	todo, err := c.Service.CreateToDoInList(&requestBody, listId)

	if err != nil {
		if err.Error() == "empty content" {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err.Error() == "list not found" {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}
	}

	ctx.JSON(http.StatusCreated, todo)

}

func (c *Controller) GetList(ctx *gin.Context) {

	var list *models.ResponseBodyList
	var err error

	listId := ctx.Param("listid")
	list, err = c.Service.GetList(listId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, list)

}

func (c *Controller) GetToDo(ctx *gin.Context) {
	todoId := ctx.Param("todoid")
	var todo *models.ToDo
	var err error

	todo, err = c.Service.GetToDoInList(todoId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (c *Controller) GetAllLists(ctx *gin.Context) {
	lists, message := c.Service.GetAllLists()

	fmt.Printf("Retrieved lists: %+v, message: %s\n", lists, message)

	if message != "" {
		ctx.JSON(http.StatusOK, message)
		return
	}

	ctx.JSON(http.StatusOK, lists)

}

func (c *Controller) GetAllTodos(ctx *gin.Context) {
	listId := ctx.Param("listid")

	todos, err := c.Service.GetAllTodos(listId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, todos)
}

func (c *Controller) PatchList(ctx *gin.Context) {
	var requestBody models.ToDoList

	if err := ctx.BindJSON(&requestBody); err != nil {
		return
	}

	listId := ctx.Param("listid")

	list, err := c.Service.PatchList(&requestBody, listId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, list)
}

func (c *Controller) PatchToDo(ctx *gin.Context) {
	var requestBody models.ToDo

	if err := ctx.BindJSON(&requestBody); err != nil {
		return
	}

	todoId := ctx.Param("todoid")

	todo, err := c.Service.PatchToDoInList(&requestBody, todoId)

	if err != nil {

		if err.Error() == "todo already completed" {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err.Error() == "todo already not completed" {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err.Error() == "todo not found" {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}

	}

	ctx.JSON(http.StatusOK, todo)
}

func (c *Controller) DeleteList(ctx *gin.Context) {
	listId := ctx.Param("listid")

	err := c.Service.DeleteList(listId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.Status(http.StatusNoContent)

}

func (c *Controller) DeleteToDo(ctx *gin.Context) {
	todoId := ctx.Param("todoid")

	err := c.Service.DeleteToDoInList(todoId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *Controller) GetDataStructure(ctx *gin.Context) {
	dataStructure, message := c.Service.GetDataStructure()

	if message != "" {
		ctx.JSON(http.StatusOK, message)
		return
	}

	ctx.JSON(http.StatusOK, dataStructure)
}
