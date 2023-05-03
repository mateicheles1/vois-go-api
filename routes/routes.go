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

	data := data.NewToDoListDB(data.ReturnDB())
	service := service.NewToDoListService(data)
	controller := controllers.NewController(service)

	r := gin.New()

	r.Use(middlewares.ErrorHandler())
	r.Use(middlewares.InfoHandler())
	r.Use(gin.Recovery())

	r.GET("api/v2/lists", controller.GetLists)
	r.GET("api/v2/lists/:listid/todos", controller.GetTodos)

	r.GET("api/v2/lists/:listid", controller.GetList)
	r.POST("api/v2/lists", controller.CreateList)
	r.PATCH("api/v2/lists/:listid", controller.PatchList)
	r.DELETE("api/v2/lists/:listid", controller.DeleteList)

	r.GET("api/v2/todos/:todoid", controller.GetTodo)
	r.POST("api/v2/lists/:listid/todos", controller.CreateTodo)
	r.PATCH("api/v2/todos/:todoid", controller.PatchTodo)
	r.DELETE("api/v2/todos/:todoid", controller.DeleteTodo)

	if err := r.Run(":8080"); err != nil {
		logs.ErrorLogger.Fatal().Msgf("Could not start the server due to: %s", err)
	}
}
