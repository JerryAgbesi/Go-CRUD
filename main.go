package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryagbesi/gocrud/Controllers"
	"github.com/jerryagbesi/gocrud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()

}

func main() {
	r := gin.Default()
	r.POST("items/", Controllers.ItemCreate)
	r.GET("items/", Controllers.ItemIndex)
	r.GET("items/:id", Controllers.ItemShow)
	r.PUT("items/:id", Controllers.ItemUpdate)
	r.DELETE("items/:id", Controllers.ItemDelete)
	r.Run() // listen and serve on 0.0.0.0:8080

}
