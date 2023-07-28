package routes

import (
	"MileTravel/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func LoadRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/api/testimonials-home", controllers.TestimonialsHome)
	r.GET("/api/testimonials", controllers.Testimonials)
	r.POST("/api/testimonials", controllers.CreateTestimonial)
	r.GET("/api/testimonials/:id", controllers.TestimonialById)
	r.DELETE("/api/testimonials/:id", controllers.DeleteTestimonial)
	r.PUT("/api/testimonials/:id", controllers.UpdateTestimonial)

	return r
}
