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
	service service.ToDoListServiceInterface
}

func NewController(service service.ToDoListServiceInterface) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) GetLists(ctx *gin.Context) {

	lists, err := c.service.GetLists()

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, lists)
}

func (c *Controller) GetTodos(ctx *gin.Context) {

	todos, err := c.service.GetTodos(ctx.Param("listid"))

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, "list not found")
			return
		}

		if err.Error() == "invalid UUID format" {
			ctx.AbortWithError(http.StatusBadRequest, err)
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

	if reqBody.Owner == "" {
		ctx.JSON(http.StatusBadRequest, "empty owner")
		return
	}

	if len(reqBody.Todos) == 0 {
		ctx.JSON(http.StatusBadRequest, "empty todos")
		return
	}

	list, err := c.service.CreateList(&reqBody)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Header("Location", fmt.Sprintf("/api/v2/lists/%s", list.Id))

	ctx.JSON(http.StatusCreated, list)
}

func (c *Controller) GetList(ctx *gin.Context) {

	list, err := c.service.GetList(ctx.Param("listid"))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, "list not found")
			return
		}

		if err.Error() == "invalid UUID format" {
			ctx.AbortWithError(http.StatusBadRequest, err)
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

	if reqBody.Owner == "" {
		ctx.JSON(http.StatusBadRequest, "empty owner")
		return
	}

	list, err := c.service.PatchList(&reqBody, ctx.Param("listid"))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, "list not found")
			return
		}

		if err.Error() == "invalid UUID format" {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, list)
}

func (c *Controller) DeleteList(ctx *gin.Context) {

	err := c.service.DeleteList(ctx.Param("listid"))

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, "list not found")
			return
		}

		if err.Error() == "invalid UUID format" {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *Controller) CreateTodo(ctx *gin.Context) {
	var reqBody models.ToDo

	if err := ctx.BindJSON(&reqBody); err != nil {
		return
	}

	if reqBody.Content == "" {
		ctx.JSON(http.StatusBadRequest, "empty content")
		return
	}

	todo, err := c.service.CreateTodo(&reqBody, ctx.Param("listid"))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, "list not found")
			return
		}

		if err.Error() == "invalid UUID format" {
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

	todo, err := c.service.GetTodo(ctx.Param("todoid"))

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, "todo not found")
			return
		}

		if err.Error() == "invalid UUID format" {
			ctx.AbortWithError(http.StatusBadRequest, err)
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

	todo, err := c.service.PatchTodo(&reqBody, ctx.Param("todoid"))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, "todo not found")
		}

		if err.Error() == "invalid UUID format" {
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

	err := c.service.DeleteTodo(todoId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, "todo not found")
			return
		}

		if err.Error() == "invalid UUID format" {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
