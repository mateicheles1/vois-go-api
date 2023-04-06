package routes

import (
	"gogin-api/controllers"
	"gogin-api/logs"
	"gogin-api/middleware"

	"github.com/gin-gonic/gin"
)


func SetupRoutes() {
	r := gin.Default()

	r.Use(middleware.ErrorHandler)

	r.GET("api/v2/lists", controllers.Lists)
	r.GET("api/v2/list/:listid/todos", controllers.Todos)
	
	r.GET("/api/v2/list/:listid", controllers.GetList)
	r.POST("/api/v2/list", controllers.CreateList)
	r.PATCH("api/v2/list/:listid", controllers.PatchList)
	r.DELETE("api/v2/list/:listid", controllers.DeleteList)

	r.GET("api/v2/list/:listid/todo/:todoid", controllers.GetToDo)
	r.POST("api/v2/list/:listid/todo", controllers.CreateToDo)
	r.PATCH("/api/v2/list/:listid/todo/:todoid", controllers.PatchToDo)
	r.DELETE("api/v2/list/:listid/todo/:todoid", controllers.DeleteToDo)

	if err := r.Run(); err != nil {
		logs.Logger.Fatal().Msgf("Could not start the server due to: %s", err.Error())
	}
}