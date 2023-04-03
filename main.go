package main

import (
	"gogin-api/initializers"
	"gogin-api/routes"
)

func init() {
	initializers.LoadEnvVariables()
	// initializers.ConnectToDB()
}
 
func main() {
	routes.SetupRoutes()
}