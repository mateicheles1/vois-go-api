package main

import (
	"gogin-api/logs"
	"gogin-api/routes"
)

func main() {
	defer logs.LogFile.Close()
	routes.SetupRoutes()
}
