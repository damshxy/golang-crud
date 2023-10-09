package handlers

import (
	"go-paslons-crud/config"
	"go-paslons-crud/models"
	"go-paslons-crud/services"

	"github.com/gin-gonic/gin"
)

func CreatePaslons(c *gin.Context) {
	name := c.PostForm("name")
	visi := c.PostForm("visi")

	formfile, _, err := c.Request.FormFile("image")
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

	uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: formfile})
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

	paslonModel := models.Paslons{
        Name:  name,
        Visi:  visi,
        Image: uploadUrl,
	}

	config.DB.Create(&paslonModel)
    c.JSON(200, paslonModel)
}

func GetPaslons(c *gin.Context) {
	var paslons []models.Paslons
	
	config.DB.Preload("Votes").Preload("Parties").Find(&paslons)

	c.JSON(200, gin.H{
		"message": "paslons data found",
		"paslons": paslons,
	})
}

func GetPaslonById(c *gin.Context) {
	id := c.Param("id")
	var paslons models.Paslons

	config.DB.Preload("Votes").Preload("Parties").First(&paslons, id)

	c.JSON(200, paslons)
}

func UpadatePaslon(c *gin.Context) {
	id := c.Param("id")
	var paslonBody struct {
		Name string
		Visi string
	}

	c.Bind(&paslonBody)

	var paslons models.Paslons
	config.DB.First(&paslons, id)

	config.DB.Model(&paslons).Updates(models.Paslons{
		Name: paslonBody.Name,
		Visi: paslonBody.Visi,
	})

	c.JSON(200, gin.H{
		"message": "Update data succesfully",
		"paslons": paslons,
	})
}

func DeletePaslon(c *gin.Context) {
	id := c.Param("id")

	config.DB.Delete(&models.Paslons{}, id)

	c.JSON(200, gin.H{"message": "Paslon Deleted !!!"})
}