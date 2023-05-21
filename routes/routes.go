package routes

import (
	"fmt"
	"gogin-api/controllers"
	"gogin-api/data"
	"gogin-api/logs"
	"gogin-api/middlewares"
	"gogin-api/service"

	"gogin-api/config"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() {

	config := config.NewConfig("./config/config.json")
	db := data.ConnectToDB(config)
	data := data.NewToDoListDB(db)
	service := service.NewToDoListService(data)
	controller := controllers.NewController(service)

	r := gin.New()

	r.Use(middlewares.ErrorHandler())
	r.Use(middlewares.InfoHandler())
	r.Use(gin.Recovery())

	authRoute := r.Group("api/v2")

	authRoute.Use(middlewares.AuthMiddleware())

	r.POST("api/v2/login", controller.Login)
	r.POST("api/v2/signup", controller.CreateUser)

	authRoute.GET("lists", controller.GetLists)
	authRoute.GET("lists/:listid/todos", controller.GetTodos)

	authRoute.GET("lists/:listid", controller.GetList)
	authRoute.POST("lists", controller.CreateList)
	authRoute.PATCH("lists/:listid", controller.PatchList)
	authRoute.DELETE("lists/:listid", controller.DeleteList)

	authRoute.GET("todos/:todoid", controller.GetTodo)
	authRoute.POST("lists/:listid/todos", controller.CreateTodo)
	authRoute.PATCH("todos/:todoid", controller.PatchTodo)
	authRoute.DELETE("todos/:todoid", controller.DeleteTodo)

	if err := r.Run(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)); err != nil {
		logs.ErrorLogger.Fatal().Msgf("Could not start the server due to: %s", err)
	}

}
