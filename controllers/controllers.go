package controllers

import (
	"gogin-api/models"
	"gogin-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service service.ToDoListService
}

func (h Handler) GetAllListsHandler(c *gin.Context) {
	lists, err := h.Service.GetAllLists()

	if err != nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, lists)
}

func (h Handler) GetListHandler(c *gin.Context) {
	list, err := h.Service.GetList(c.Param("listid"))

	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) CreateListHandler(c *gin.Context) {
	requestBody := new(models.RequestBodyList)

	if err := c.BindJSON(requestBody); err != nil {
		return
	}

	h.Service.CreateList(requestBody)

	c.Status(http.StatusCreated)
}

func (h *Handler) PatchListHandler(c *gin.Context) {
	requestBody := new(models.ToDoList)

	if err := c.BindJSON(requestBody); err != nil {
		return
	}

	requestBody.Id = c.Param("listid")

	err := h.Service.PatchList(requestBody)

	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteListHandler(c *gin.Context) {
	err := h.Service.DeleteList(c.Param("listid"))

	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK)
}

func (h Handler) GetToDoHandler(c *gin.Context) {
	todo, err := h.Service.GetToDoInList(c.Param("todoid"))

	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *Handler) DeleteToDoHandler(c *gin.Context) {
	err := h.Service.DeleteToDoInList(c.Param("todoid"))

	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) PatchToDoHandler(c *gin.Context) {
	requestBody := new(models.ToDo)

	if err := c.BindJSON(requestBody); err != nil {
		return
	}

	err := h.Service.PatchToDoInList(requestBody.Completed, c.Param("todoid"))

	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) CreateToDoHandler(c *gin.Context) {
	requestBody := new(models.ToDo)

	if err := c.BindJSON(requestBody); err != nil {
		return
	}
	if requestBody.Content == "" {
		c.String(http.StatusBadRequest, "content can't be empty")
		return
	}
	err := h.Service.CreateToDoInList(c.Param("listid"), requestBody.Content)

	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusCreated)
}

func (h Handler) GetAllToDosHandler(c *gin.Context) {
	todos, err := h.Service.GetAllToDosInList(c.Param("listid"))

	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, todos)
}

// func GetAllListsHandler(todoListService service.ToDoListService) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		lists, err := todoListService.GetAllLists()
// 		if err != nil {
// 			c.Status(204)
// 			return
// 		}
// 		c.JSON(http.StatusOK, lists)
// 	}

// }

// func GetAllTodosHandler(todoListService service.ToDoListService) gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		todos, err := todoListService.GetAllToDosInList(c.Param("listid"))

// 		if err != nil {
// 			c.Status(http.StatusNotFound)
// 			return
// 		}

// 		c.JSON(http.StatusOK, todos)
// 	}

// }

// func GetListHandler(todoListService service.ToDoListService) gin.HandlerFunc {

// 	return func(c *gin.Context) {

// 		list, err := todoListService.GetList(c.Param("listid"))

// 		if err != nil {
// 			c.Status(http.StatusNotFound)
// 			return
// 		}

// 		c.JSON(http.StatusOK, list)
// 	}

// }

// func CreateListHandler(todoListService service.ToDoListService) gin.HandlerFunc {

// 	return func(c *gin.Context) {

// 		requestBody := new(models.RequestBodyList)

// 		if err := c.BindJSON(requestBody); err != nil {
// 			return
// 		}

// 		todoListService.CreateList(requestBody)

// 		c.Status(http.StatusCreated)
// 	}

// }

// func PatchListHandler(todoListService service.ToDoListService) gin.HandlerFunc {

// 	return func(c *gin.Context) {

// 		requestBody := new(models.ToDoList)

// 		if err := c.BindJSON(requestBody); err != nil {
// 			return
// 		}

// 		requestBody.Id = c.Param("listid")

// 		err := todoListService.PatchList(requestBody)

// 		if err != nil {
// 			c.Status(http.StatusNotFound)
// 			return
// 		}

// 		c.Status(http.StatusOK)
// 	}

// }

// func DeleteListHandler(todoListService service.ToDoListService) gin.HandlerFunc {

// 	return func(c *gin.Context) {

// 		err := todoListService.DeleteList(c.Param("listid"))

// 		if err != nil {
// 			c.Status(http.StatusNotFound)
// 			return
// 		}

// 		c.Status(http.StatusOK)
// 	}

// }

// func GetToDoHandler(todoListService service.ToDoListService) gin.HandlerFunc {

// 	return func(c *gin.Context) {

// 		todo, err := todoListService.GetToDoInList(c.Param("todoid"))

// 		if err != nil {
// 			c.Status(http.StatusNotFound)
// 			return
// 		}

// 		c.JSON(http.StatusOK, todo)
// 	}

// }

// func DeleteToDoHandler(todoListService service.ToDoListService) gin.HandlerFunc {

// 	return func(c *gin.Context) {

// 		err := todoListService.DeleteToDoInList(c.Param("todoid"))

// 		if err != nil {
// 			c.Status(http.StatusNotFound)
// 			return
// 		}

// 		c.Status(http.StatusOK)
// 	}

// }

// func PatchToDoHandler(todoListService service.ToDoListService) gin.HandlerFunc {

// 	return func(c *gin.Context) {

// 		requestBody := new(models.ToDo)

// 		if err := c.BindJSON(requestBody); err != nil {
// 			return
// 		}

// 		err := todoListService.PatchToDoInList(requestBody.Completed, c.Param("todoid"))

// 		if err != nil {
// 			c.Status(http.StatusNotFound)
// 			return
// 		}

// 		c.Status(http.StatusOK)
// 	}
// }

// func CreateToDoHandler(todoListService service.ToDoListService) gin.HandlerFunc {

// 	return func(c *gin.Context) {

// 		requestBody := new(models.ToDo)

// 		if err := c.BindJSON(requestBody); err != nil {
// 			return
// 		}
// 		if requestBody.Content == "" {
// 			c.String(http.StatusBadRequest, "content can't be empty")
// 			return
// 		}
// 		err := todoListService.CreateToDoInList(c.Param("listid"), requestBody.Content)

// 		if err != nil {
// 			c.Status(http.StatusNotFound)
// 			return
// 		}

// 		c.Status(http.StatusCreated)
// 	}

// }
