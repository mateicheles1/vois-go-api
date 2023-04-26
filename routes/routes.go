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

	// controller-ul care primeste o implementare a interfetei care la randul ei, tot prin constructor function, primeste un struct de tip todolistdb. am ales sa fac asa ca sa pot schimba implementarea interfetei si orice instantare a struct-ului `ToDoListDB`, astfel utilizand dependency injection si loose coupling a diverselor componente din app.

	controller := controllers.NewController(service.NewToDoListService(data.ToDoListDB{}))

	r := gin.New()

	r.Use(middlewares.ErrorHandler())
	r.Use(middlewares.InfoHandler())
	r.Use(gin.Recovery())

	r.GET("api/v2/lists", controller.GetAllListsController)
	r.GET("api/v2/lists/:listid/todos", controller.GetAllToDosController)

	r.GET("api/v2/lists/:listid", controller.GetListController)
	r.POST("api/v2/lists", controller.CreateListController)
	r.PATCH("api/v2/lists/:listid", controller.PatchListController)
	r.DELETE("api/v2/lists/:listid", controller.DeleteListController)

	r.GET("api/v2/todos/:todoid", controller.GetToDoController)
	r.POST("api/v2/lists/:listid/todos", controller.CreateToDoController)
	r.PATCH("api/v2/todos/:todoid", controller.PatchToDoController)
	r.DELETE("api/v2/todos/:todoid", controller.DeleteToDoController)

	// ruta sa vad intreaga structura de date. nu face parte din api
	r.GET("api/v2/data-structure", controller.GetDataStructureController)

	if err := r.Run(); err != nil {
		logs.ErrorLogger.Fatal().Msgf("Could not start the server due to: %s", err.Error())
	}
}
