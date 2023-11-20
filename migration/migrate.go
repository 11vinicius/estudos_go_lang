package main

import (
	"boilerplate/initializers"
	"boilerplate/models"
)

func main() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&models.Users{})
}
