package main

import (
	"github.com/jerryagbesi/gocrud/initializers"
	"github.com/jerryagbesi/gocrud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()

}

func main() {
	initializers.DB.AutoMigrate(&models.Item{})
}
