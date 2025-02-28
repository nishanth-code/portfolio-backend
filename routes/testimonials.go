package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nishanth-code/portfolio-backend/auth"
	"github.com/nishanth-code/portfolio-backend/controllers"
)

func RegisterTestimonialRoutes(router *gin.RouterGroup){
	testimonialRoutes := router.Group("/testimonials")
	
	testimonialRoutes.GET("/",controllers.GetTestimonial)
	testimonialRoutes.POST("/",auth.AuthMiddleware(),controllers.InsertTestimonial)
	testimonialRoutes.PATCH("/:id",auth.AuthMiddleware(),controllers.UpdateTestimonial)
	testimonialRoutes.DELETE("/:id",auth.AuthMiddleware(),controllers.DeleteTestimonial)

}