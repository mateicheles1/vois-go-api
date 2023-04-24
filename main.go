package main

import (
	"gogin-api/initializers"
	"gogin-api/logs"
	"gogin-api/routes"
)

func init() {
	initializers.LoadEnvVariables()
	// initializers.ConnectToDB()
}

func main() {
	defer logs.LogFile.Close()
	routes.SetupRoutes()
}
