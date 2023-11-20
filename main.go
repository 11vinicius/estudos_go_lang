package main

import (
	"boilerplate/initializers"
	"boilerplate/routes"
)

func main() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	routes.Index()
}
