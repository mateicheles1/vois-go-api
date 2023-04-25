package routes

import (
	"gogin-api/controllers"
	"gogin-api/data"
	"gogin-api/logs"
	"gogin-api/middlewares"
	"gogin-api/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() {

	// controller made from 2 constructor function. 1. the constructor function that instances the service and 2. the const function that instances the handler

	controller := controllers.NewController(service.NewToDoListService(data.ToDoListDB{}))

	r := gin.New()

	r.Use(middlewares.ErrorHandler())
	r.Use(middlewares.InfoHandler())
	r.Use(gin.Recovery())

	r.GET("api/v2/lists", controller.CreateListController)
	r.GET("api/v2/lists/:listid/todos", controller.GetAllToDosController)

	r.GET("api/v2/lists/:listid", controller.GetListController)
	r.POST("api/v2/lists", controller.CreateListController)
	r.PATCH("api/v2/lists/:listid", controller.PatchListController)
	r.DELETE("api/v2/lists/:listid", controller.DeleteListController)

	r.GET("api/v2/todos/:todoid", controller.GetToDoController)
	r.POST("api/v2/lists/:listid/todos", controller.CreateToDoController)
	r.PATCH("api/v2/todos/:todoid", controller.PatchToDoController)
	r.DELETE("api/v2/todos/:todoid", controller.DeleteToDoController)

	// route to check the entire data structure. isn't part of the api
	r.GET("api/v2/data-structure", controller.GetDataStructureController)

	if err := r.Run(); err != nil {
		logs.ErrorLogger.Fatal().Msgf("Could not start the server due to: %s", err.Error())
	}
}
