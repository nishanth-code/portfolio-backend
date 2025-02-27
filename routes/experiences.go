package routes

import("github.com/gin-gonic/gin")

func RegisterExperienceRoutes(router *gin.RouterGroup){
	experienceRoutes := router.Group("/experience")
	
	experienceRoutes.GET("/",)
	experienceRoutes.POST("/",)
	experienceRoutes.PATCH("/:id",)
	experienceRoutes.DELETE("/",)

}