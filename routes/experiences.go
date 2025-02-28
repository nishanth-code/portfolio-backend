package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nishanth-code/portfolio-backend/auth"
	"github.com/nishanth-code/portfolio-backend/controllers"
)

func RegisterExperienceRoutes(router *gin.RouterGroup){
	experienceRoutes := router.Group("/experience")
	
	experienceRoutes.GET("/",controllers.GetExperience)
	experienceRoutes.POST("/",auth.AuthMiddleware(),controllers.InsertExperience)
	experienceRoutes.PATCH("/:id",auth.AuthMiddleware(),controllers.UpdateExperience)
	experienceRoutes.DELETE("/:id",auth.AuthMiddleware(),controllers.DeleteExperience)

}