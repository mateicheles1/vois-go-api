package routes

import (
	"gogin-api/controllers"
	"gogin-api/logs"
	"gogin-api/middlewares"
	"gogin-api/service"

	"github.com/gin-gonic/gin"
)

func newHandler(service service.ToDoListServiceInterface) controllers.Handler {
	return controllers.Handler{
		Service: service,
	}
}

func SetupRoutes() {
	r := gin.New()

	handler := newHandler(&service.ToDoListService{})

	r.Use(middlewares.ErrorHandler())
	r.Use(middlewares.InfoHandler())
	r.Use(gin.Recovery())

	r.GET("api/v2/lists", handler.GetAllListsHandler)
	r.GET("api/v2/lists/:listid/todos", handler.GetAllToDosHandler)

	r.GET("api/v2/lists/:listid", handler.GetListHandler)
	r.POST("api/v2/lists", handler.CreateListHandler)
	r.PATCH("api/v2/lists/:listid", handler.PatchListHandler)
	r.DELETE("api/v2/lists/:listid", handler.DeleteListHandler)

	r.GET("api/v2/todos/:todoid", handler.GetToDoHandler)
	r.POST("api/v2/lists/:listid/todos", handler.CreateToDoHandler)
	r.PATCH("api/v2/todos/:todoid", handler.PatchToDoHandler)
	r.DELETE("api/v2/todos/:todoid", handler.DeleteToDoHandler)

	r.GET("api/v2/data-structure", handler.GetDataStructureHandler)

	if err := r.Run(); err != nil {
		logs.ErrorLogger.Fatal().Msgf("Could not start the server due to: %s", err.Error())
	}
}
