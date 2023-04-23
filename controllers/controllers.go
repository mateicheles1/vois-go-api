package controllers

import (
	"gogin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		c.JSON(http.StatusNotFound, err.Error())
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
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteListHandler(c *gin.Context) {
	err := h.Service.DeleteList(c.Param("listid"))

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h Handler) GetToDoHandler(c *gin.Context) {
	todo, err := h.Service.GetToDoInList(c.Param("todoid"))

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *Handler) DeleteToDoHandler(c *gin.Context) {
	err := h.Service.DeleteToDoInList(c.Param("todoid"))

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
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
		c.JSON(http.StatusNotFound, err.Error())
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
	if err := h.Service.CreateToDoInList(c.Param("listid"), requestBody.Content); err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

func (h Handler) GetAllToDosHandler(c *gin.Context) {
	todos, err := h.Service.GetAllToDosInList(c.Param("listid"))

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (h Handler) GetDataStructureHandler(c *gin.Context) {
	dataStructure, err := h.Service.GetDataStructure()

	if err != nil {
		c.JSON(http.StatusNoContent, err.Error())
		return
	}

	c.JSON(200, dataStructure)
}
