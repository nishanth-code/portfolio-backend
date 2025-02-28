package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nishanth-code/portfolio-backend/controllers"
)

func RegisterTestimonialRoutes(router *gin.RouterGroup){
	testimonialRoutes := router.Group("/testimonials")
	
	testimonialRoutes.GET("/",controllers.GetTestimonial)
	testimonialRoutes.POST("/",controllers.InsertTestimonial)
	testimonialRoutes.PATCH("/:id",controllers.UpdateTestimonial)
	testimonialRoutes.DELETE("/:id",controllers.DeleteTestimonial)

}