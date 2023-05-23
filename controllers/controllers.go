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

var ErrInvalidUUID service.UuidError

func (c *Controller) GetLists(ctx *gin.Context) {

	username := ctx.MustGet("username").(string)
	role := ctx.MustGet("role").(string)

	lists, err := c.service.GetLists(username, role)

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

		if errors.As(err, &ErrInvalidUUID) {
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

	owner := ctx.MustGet("username").(string)

	reqBody.Owner = owner

	if len(reqBody.Todos) == 0 {
		ctx.JSON(http.StatusBadRequest, "empty todos")
		return
	}

	list, err := c.service.CreateList(&reqBody, reqBody.Owner)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Header("Location", fmt.Sprintf("/api/v2/lists/%s", list.Id))

	ctx.JSON(http.StatusCreated, list)
}

func (c *Controller) GetList(ctx *gin.Context) {

	username := ctx.MustGet("username").(string)
	role := ctx.MustGet("role").(string)

	list, err := c.service.GetList(ctx.Param("listid"), username, role)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, "list not found")
			return
		}

		if errors.As(err, &ErrInvalidUUID) {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if err.Error() == "action not allowed" {
			ctx.AbortWithError(http.StatusForbidden, err)
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

		if errors.As(err, &ErrInvalidUUID) {
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

		if errors.As(err, &ErrInvalidUUID) {
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

		if errors.As(err, &ErrInvalidUUID) {
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

		if errors.As(err, &ErrInvalidUUID) {
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
			return
		}

		if errors.As(err, &ErrInvalidUUID) {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (c *Controller) DeleteTodo(ctx *gin.Context) {

	err := c.service.DeleteTodo(ctx.Param("todoid"))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, "todo not found")
			return
		}

		if errors.As(err, &ErrInvalidUUID) {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *Controller) DeleteAllLists(ctx *gin.Context) {

	role := ctx.MustGet("role").(string)

	err := c.service.DeleteAllLists(role)

	if err != nil {

		if err.Error() == "action not allowed" {
			ctx.AbortWithError(http.StatusForbidden, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *Controller) CreateUser(ctx *gin.Context) {
	var reqBody models.User

	if err := ctx.BindJSON(&reqBody); err != nil {
		return
	}

	if reqBody.Role == "" {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("empty owner"))
		return
	}

	user, err := c.service.CreateUser(&reqBody)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, user)

}

func (c *Controller) Login(ctx *gin.Context) {
	var reqBody models.User

	if err := ctx.BindJSON(&reqBody); err != nil {
		return
	}

	token, err := c.service.Login(&reqBody)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, token)
}

func (c *Controller) DeleteUser(ctx *gin.Context) {

	role := ctx.MustGet("role").(string)

	err := c.service.DeleteUser(ctx.Param("username"), role)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, "user not found")
			return
		}

		if err.Error() == "action not allowed" {
			ctx.AbortWithError(http.StatusForbidden, err)
			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
