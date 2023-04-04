package routes

import (
	"gogin-api/controllers"
	"gogin-api/logs"

	"github.com/gin-gonic/gin"
)


func SetupRoutes() {
	r := gin.Default()
	listRouter := r.Group("/api/v2/list")
	todoRouter := r.Group("/api/v2/list/:listid/todo")

	r.GET("api/v2/lists", controllers.Lists)
	r.GET("api/v2/list/:listid/todos", controllers.Todos)
	
	listRouter.GET("/:listid", controllers.GetList)
	listRouter.POST("/", controllers.CreateList)
	listRouter.PATCH("/:listid", controllers.PatchList)
	listRouter.DELETE("/:listid", controllers.DeleteList)

	todoRouter.GET("/:todoid", controllers.GetToDo)
	todoRouter.POST("/", controllers.CreateToDo)
	todoRouter.PATCH("/:todoid", controllers.PatchToDo)
	todoRouter.DELETE("/:todoid", controllers.DeleteToDo)

	if err := r.Run(); err != nil {
		logs.Logger.Fatal().Msgf("Could not start the server due to: %s", err.Error())
	}
}