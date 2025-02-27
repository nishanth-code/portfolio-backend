package routes

import("github.com/gin-gonic/gin")

func RegisterTestimonialRoutes(router *gin.RouterGroup){
	testimonialRoutes := router.Group("/testimonials")
	
	testimonialRoutes.GET("/",)
	testimonialRoutes.POST("/",)
	testimonialRoutes.PATCH("/:id",)
	testimonialRoutes.DELETE("/",)

}