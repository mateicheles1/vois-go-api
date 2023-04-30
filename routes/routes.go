package routes

import (
	"gogin-api/controllers"
	"gogin-api/data"
	"gogin-api/logs"
	"gogin-api/middlewares"
	"gogin-api/models"
	"gogin-api/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() {

	lists := make(map[string]*models.ToDoList)
	data := data.NewToDoListDB(lists)
	service := service.NewToDoListService(data)
	controller := controllers.NewController(service)

	r := gin.New()

	r.Use(middlewares.ErrorHandler())
	r.Use(middlewares.InfoHandler())
	r.Use(gin.Recovery())

	r.GET("api/v2/lists", controller.GetAllLists)
	r.GET("api/v2/lists/:listid/todos", controller.GetAllTodos)

	r.GET("api/v2/lists/:listid", controller.GetList)
	r.POST("api/v2/lists", controller.CreateList)
	r.PATCH("api/v2/lists/:listid", controller.PatchList)
	r.DELETE("api/v2/lists/:listid", controller.DeleteList)

	r.GET("api/v2/todos/:todoid", controller.GetToDo)
	r.POST("api/v2/lists/:listid/todos", controller.CreateToDo)
	r.PATCH("api/v2/todos/:todoid", controller.PatchToDo)
	r.DELETE("api/v2/todos/:todoid", controller.DeleteToDo)

	// ruta sa vad intreaga structura de date. nu face parte din api
	r.GET("api/v2/data-structure", controller.GetDataStructure)

	if err := r.Run(); err != nil {
		logs.ErrorLogger.Fatal().Msgf("Could not start the server due to: %s", err)
	}
}
