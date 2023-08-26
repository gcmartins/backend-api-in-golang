package controllers

import (
	"MileTravel/database"
	"MileTravel/models"
	"MileTravel/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TestimonialsHome(c *gin.Context) {
	testimonials := []models.Testimonial{}

	sql := database.DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&models.Testimonial{}).Order("random()").Limit(3).Find(&testimonials)
	})

	database.DB.Raw(sql).Scan(&testimonials)
	c.JSON(200, testimonials)
}

func Testimonials(c *gin.Context) {
	var testimonials []models.Testimonial
	database.DB.Find(&testimonials)

	c.JSON(http.StatusOK, testimonials)

}

func TestimonialById(c *gin.Context) {
	id := c.Params.ByName("id")
	var testimonial models.Testimonial
	database.DB.First(&testimonial, id)
	if testimonial.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Resource not found."})
		return
	}
	c.JSON(http.StatusOK, testimonial)
}

func CreateTestimonial(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	uploadsPath := storage.GetStoragePath()

	fmt.Println(uploadsPath)

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filePath := filepath.Join(uploadsPath, file.Filename)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var testimonial models.Testimonial

	testimonialJson := c.Request.FormValue("json")
	err = json.Unmarshal([]byte(testimonialJson), &testimonial)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&testimonial)
	c.JSON(http.StatusOK, testimonial)
}

func DeleteTestimonial(c *gin.Context) {
	id := c.Params.ByName("id")
	var testimonial models.Testimonial
	database.DB.Delete(&testimonial, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Resource deleted."})
}

func UpdateTestimonial(c *gin.Context) {
	id := c.Params.ByName("id")
	var testimonial models.Testimonial
	database.DB.First(&testimonial, id)

	if testimonial.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Resource not found."})
		return
	}

	if err := c.ShouldBindJSON(&testimonial); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Save(&testimonial)
	c.JSON(http.StatusOK, testimonial)
}
