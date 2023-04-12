package routes

import (
	"gogin-api/controllers"
	"gogin-api/logs"
	"gogin-api/middlewares"

	"github.com/gin-gonic/gin"
)


func SetupRoutes() {
	r := gin.New()

	r.Use(middlewares.ErrorHandler)
	r.Use(middlewares.InfoHandler())

	r.GET("api/v2/lists", controllers.Lists)
	r.GET("api/v2/lists/:listid/todos", controllers.Todos)
	
	r.GET("/api/v2/lists/:listid", controllers.GetList)
	r.POST("/api/v2/lists", controllers.CreateList)
	r.PATCH("api/v2/lists/:listid", controllers.PatchList)
	r.DELETE("api/v2/lists/:listid", controllers.DeleteList)

	r.GET("api/v2/todos/:todoid", controllers.GetToDo)
	r.POST("api/v2/lists/:listid/todos", controllers.CreateToDo)
	r.PATCH("/api/v2/todos/:todoid", controllers.PatchToDo)
	r.DELETE("api/v2//todos/:todoid", controllers.DeleteToDo)

	if err := r.Run(); err != nil {
		logs.ErrorLogger.Fatal().Msgf("Could not start the server due to: %s", err.Error())
	}
}