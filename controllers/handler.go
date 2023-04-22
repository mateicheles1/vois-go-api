package controllers

import "gogin-api/service"

type Handler struct {
	Service service.ToDoListServiceInterface
}

func NewHandler(service service.ToDoListServiceInterface) Handler {
	return Handler{
		Service: service,
	}
}
