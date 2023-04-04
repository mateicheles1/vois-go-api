package routes

import (
	"gogin-api/controllers"
	"gogin-api/logs"

	"github.com/gin-gonic/gin"
)


func SetupRoutes() {
	r := gin.Default()
	r2 := r.Group("/api/v2/list/:listid/todo")

	r.GET("api/v2/lists", controllers.Lists)
	r.GET("api/v2/list/:listid/todos", controllers.Todos)
	
	r.GET("api/v2/list/:listid", controllers.GetList)
	r.POST("api/v2/list/", controllers.CreateList)
	r.PATCH("api/v2/list/:listid", controllers.PatchList)
	r.DELETE("api/v2/list/:listid", controllers.DeleteList)

	r2.GET("/:todoid", controllers.GetToDo)
	r2.POST("/", controllers.CreateToDo)
	r2.PATCH("/:todoid", controllers.PatchToDo)
	r2.DELETE("/:todoid", controllers.DeleteToDo)

	if err := r.Run(); err != nil {
		logs.Logger.Fatal().Msgf("Could not start the server due to: %s", err.Error())
	}
}