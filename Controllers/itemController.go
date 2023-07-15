package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryagbesi/gocrud/initializers"
	"github.com/jerryagbesi/gocrud/models"
)

func ItemCreate(c *gin.Context) {

	var body struct {
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Quantity int     `json:"quantity"`
	}

	c.Bind(&body)

	item := models.Item{Name: body.Name, Price: body.Price, Quantity: body.Quantity}

	result := initializers.DB.Create(&item)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"item": item,
	})
}

func ItemIndex(c *gin.Context) {
	var items []models.Item

	initializers.DB.Find(&items)

	c.JSON(200, gin.H{
		"items": items,
	})
}

func ItemShow(c *gin.Context) {
	var item models.Item

	result := initializers.DB.First(&item, c.Param("id"))


	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"item": item,
	})

}

func ItemUpdate(c *gin.Context){
	var item models.Item

	result := initializers.DB.First(&item, c.Param("id"))

	if result.Error != nil {
		c.Status(400)
		return
	}

	var body struct {
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Quantity int     `json:"quantity"`
	}

	c.Bind(&body)

	item.Name = body.Name
	item.Price = body.Price
	item.Quantity = body.Quantity

	initializers.DB.Save(&item)

	c.JSON(200, gin.H{
		"item": item,
	})

}

func ItemDelete(c *gin.Context){
	var item models.Item

	result := initializers.DB.First(&item, c.Param("id"))

	if result.Error != nil {
		c.Status(400)
		return
	}

	initializers.DB.Delete(&item)

	c.Status(203)

}