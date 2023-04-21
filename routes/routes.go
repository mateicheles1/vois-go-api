package routes

import (
	"gogin-api/controllers"
	"gogin-api/logs"
	"gogin-api/middlewares"
	"gogin-api/models"
	"gogin-api/service"

	"github.com/gin-gonic/gin"
)

func newToDoListRepo() models.ToDoListService {
	return &service.ToDoListRepo{}
}

func SetupRoutes() {
	r := gin.New()

	repo := newToDoListRepo()

	r.Use(middlewares.ErrorHandler())
	r.Use(middlewares.InfoHandler())
	r.Use(gin.Recovery())

	r.GET("api/v2/lists", controllers.GetAllListsHandler(repo))
	r.GET("api/v2/lists/:listid/todos", controllers.GetAllTodosHandler(repo))

	r.GET("/api/v2/lists/:listid", controllers.GetListHandler(repo))
	r.POST("/api/v2/lists", controllers.CreateListHandler(repo))
	r.PATCH("api/v2/lists/:listid", controllers.PatchListHandler(repo))
	r.DELETE("api/v2/lists/:listid", controllers.DeleteListHandler(repo))

	r.GET("api/v2/todos/:todoid", controllers.GetToDoHandler(repo))
	r.POST("api/v2/lists/:listid/todos", controllers.CreateToDoHandler(repo))
	r.PATCH("/api/v2/todos/:todoid", controllers.PatchToDoHandler(repo))
	r.DELETE("api/v2//todos/:todoid", controllers.DeleteToDoHandler(repo))

	// route for getting the entire data structure

	// r.GET("/api/v2/data-structure", func(c *gin.Context) {
	// 	if service.Repo == nil {
	// 		c.Status(204)
	// 	} else {
	// 		c.JSON(200, service.Repo)
	// 	}
	// })

	if err := r.Run(); err != nil {
		logs.ErrorLogger.Fatal().Msgf("Could not start the server due to: %s", err.Error())
	}
}
