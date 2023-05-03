package controllers

import (
	"errors"
	"fmt"
	"gogin-api/models"
	"gogin-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Service service.ToDoListServiceInterface
}

func NewController(service service.ToDoListServiceInterface) *Controller {
	return &Controller{
		Service: service,
	}
}

func (c *Controller) GetLists(ctx *gin.Context) {
	lists, err := c.Service.GetLists()

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, lists)
}

func (c *Controller) GetTodos(ctx *gin.Context) {

	listId := ctx.Param("listid")

	todos, err := c.Service.GetTodos(listId)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, "list not found")
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, todos)

}

func (c *Controller) CreateList(ctx *gin.Context) {
	var reqBody models.RequestBodyList

	if err := ctx.BindJSON(&reqBody); err != nil {
		return
	}

	list, err := c.Service.CreateList(&reqBody)

	if err != nil {
		if err.Error() == "empty owner" {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Header("Location", fmt.Sprintf("/api/v2/lists/%s", list.Id))

	ctx.JSON(http.StatusCreated, list)
}

func (c *Controller) GetList(ctx *gin.Context) {
	listId := ctx.Param("listid")

	list, err := c.Service.GetList(listId)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, "list not found")
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, list)
}

func (c *Controller) PatchList(ctx *gin.Context) {
	var reqBody models.RequestBodyList

	if err := ctx.BindJSON(&reqBody); err != nil {
		return
	}

	listId := ctx.Param("listid")

	list, err := c.Service.PatchList(&reqBody, listId)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, "list not found")
			return
		}

		if err.Error() == "empty owner" {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, list)
}

func (c *Controller) DeleteList(ctx *gin.Context) {
	listId := ctx.Param("listid")

	err := c.Service.DeleteList(listId)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, "list not found")
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.Status(http.StatusNoContent)
}

func (c *Controller) CreateTodo(ctx *gin.Context) {
	listId := ctx.Param("listid")
	var reqBody models.ToDo

	if err := ctx.BindJSON(&reqBody); err != nil {
		return
	}

	todo, err := c.Service.CreateTodo(&reqBody, listId)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, "list not found")
			return
		}

		if err.Error() == "empty content" {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Header("Location", fmt.Sprintf("api/v2/todos/%s", todo.Id))

	ctx.JSON(http.StatusCreated, todo)
}

func (c *Controller) GetTodo(ctx *gin.Context) {
	todoId := ctx.Param("todoid")

	todo, err := c.Service.GetTodo(todoId)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, "todo not found")
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (c *Controller) PatchTodo(ctx *gin.Context) {

	var reqBody models.ToDo

	if err := ctx.BindJSON(&reqBody); err != nil {
		return
	}

	todoId := ctx.Param("todoid")

	todo, err := c.Service.PatchTodo(&reqBody, todoId)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, "todo not found")
		}

		if err.Error() == "todo already completed" {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (c *Controller) DeleteTodo(ctx *gin.Context) {
	todoId := ctx.Param("todoid")

	err := c.Service.DeleteTodo(todoId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, "todo not found")
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
