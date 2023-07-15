package main

import (

	"github.com/gin-gonic/gin"
	"github.com/jerryagbesi/gocrud/initializers"

)


func init(){
	initializers.LoadEnvVariables()
	
}

func main(){
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to your marketlist",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}