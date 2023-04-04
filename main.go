package main

import (
	"gogin-api/initializers"
	"gogin-api/routes"
)

func init() {
	initializers.LoadEnvVariables()
	// initializers.ConnectToDB()
	// nu am cum sa termin db pentru ca nu am admin rights si imi tot arunca erori. maine o sa continui de pe laptopul personal
}
 
func main() {
	routes.SetupRoutes()
}