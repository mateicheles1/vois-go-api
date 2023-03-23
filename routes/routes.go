package routes

import (
	"gogin-api/logs"

	"github.com/gin-gonic/gin"
)


func SetupRoutes() {
	r := gin.Default()

	r.GET("api/v2/lists", lists)
	r.GET("api/v2/list/:listid/todos", todos)
	
	r.GET("api/v2/list/:listid", getList)
	r.POST("api/v2/list/", createList)
	r.PATCH("api/v2/list/:listid", updateList)
	r.DELETE("api/v2/list/:listid", deleteList)

	r.GET("api/v2/list/:listid/todo/:todoid", getToDo)
	r.POST("api/v2/list/:listid/todo", createToDo)
	r.PATCH("api/v2/list/:listid/todo/:todoid", updateToDo)
	r.DELETE("api/v2/list/:listid/todo/:todoid", deleteToDo)

	if err := r.Run(":8080"); err != nil {
		logs.Logger.Fatal().Msgf("could not start server due to: %s", err)
	}
}