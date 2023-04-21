package routes

import (
	"gogin-api/controllers"
	"gogin-api/logs"
	"gogin-api/middlewares"
	"gogin-api/service"

	"github.com/gin-gonic/gin"
)

func newHandler() *controllers.Handler {
	return &controllers.Handler{
		TodoListService: &service.ToDoListRepo{},
	}
}

func SetupRoutes() {
	r := gin.New()

	handler := newHandler()

	r.Use(middlewares.ErrorHandler())
	r.Use(middlewares.InfoHandler())
	r.Use(gin.Recovery())

	r.GET("api/v2/lists", handler.GetAllListsHandler)
	r.GET("api/v2/lists/:listid/todos", handler.GetAllToDosHandler)

	r.GET("/api/v2/lists/:listid", handler.GetListHandler)
	r.POST("/api/v2/lists", handler.CreateListHandler)
	r.PATCH("api/v2/lists/:listid", handler.PatchListHandler)
	r.DELETE("api/v2/lists/:listid", handler.DeleteListHandler)

	r.GET("api/v2/todos/:todoid", handler.GetToDoHandler)
	r.POST("api/v2/lists/:listid/todos", handler.CreateToDoHandler)
	r.PATCH("/api/v2/todos/:todoid", handler.PatchToDoHandler)
	r.DELETE("api/v2//todos/:todoid", handler.DeleteToDoHandler)

	// route for getting the entire data structure

	r.GET("/api/v2/data-structure", func(c *gin.Context) {
		if handler == nil {
			c.Status(204)
		} else {
			c.JSON(200, handler)
		}
	})

	if err := r.Run(); err != nil {
		logs.ErrorLogger.Fatal().Msgf("Could not start the server due to: %s", err.Error())
	}
}
