package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nishanth-code/portfolio-backend/auth"
	"github.com/nishanth-code/portfolio-backend/controllers"
)

func RegisterProjectRoutes(router *gin.RouterGroup){
	projectRoutes := router.Group("/projects")
	
	projectRoutes.GET("/",controllers.GetProjects)
	projectRoutes.POST("/",auth.AuthMiddleware(),controllers.InsertProject)
	projectRoutes.PATCH("/:id",auth.AuthMiddleware(),controllers.UpdateProject)
	projectRoutes.DELETE("/:id",auth.AuthMiddleware(),controllers.DeleteProject)

}