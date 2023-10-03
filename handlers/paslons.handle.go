package handlers

import (
	"go-paslons-crud/config"
	"go-paslons-crud/models"

	"github.com/gin-gonic/gin"
)

func CreatePaslons(c *gin.Context) {
	var paslonCreate models.Paslons

	if err := c.ShouldBindJSON(&paslonCreate)
	err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
	}

	paslonModel := models.Paslons{
		Name: paslonCreate.Name,
		Visi: paslonCreate.Visi,
		Image: paslonCreate.Image,
	}

	config.DB.Create(&paslonModel)

	c.JSON(200, paslonModel)
}

func GetPaslons(c *gin.Context) {
	var paslons []models.Paslons
	
	if err := config.DB.Find(&paslons)

	err.Error != nil {
		c.JSON(404, gin.H{"Error": "Paslons not found !!!"})
	}

	c.JSON(200, paslons)
}

func GetPaslonById(c *gin.Context) {
	id := c.Param("id")
	var paslons models.Paslons

	config.DB.First(&paslons, id)

	c.JSON(200, paslons)
}

func UpadatePaslon(c *gin.Context) {
	id := c.Param("id")
	var paslonBody struct {
		Name string
		Visi string
		Image string
	}

	c.Bind(&paslonBody)

	var paslons models.Paslons
	
	config.DB.First(&paslons, id)

	config.DB.Model(&paslons).Updates(models.Paslons{
		Name: paslonBody.Name,
		Visi: paslonBody.Visi,
		Image: paslonBody.Image,
	})

	c.JSON(200, paslons)
}

func DeletePaslon(c *gin.Context) {
	id := c.Param("id")

	config.DB.Delete(&models.Paslons{}, id)

	c.JSON(200, gin.H{"message": "Paslon Deleted !!!"})
}