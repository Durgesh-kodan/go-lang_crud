package main

import (
	"example.com/golang-crud/initializers"
	model "example.com/golang-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&model.Post{})

}
