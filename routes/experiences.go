package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nishanth-code/portfolio-backend/controllers"
)

func RegisterExperienceRoutes(router *gin.RouterGroup){
	experienceRoutes := router.Group("/experience")
	
	experienceRoutes.GET("/",controllers.GetPExperience)
	experienceRoutes.POST("/",controllers.InsertExperience)
	experienceRoutes.PATCH("/:id",controllers.UpdateExperience)
	experienceRoutes.DELETE("/:id",controllers.DeleteExperience)

}