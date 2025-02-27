package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nishanth-code/portfolio-backend/controllers"
)

func RegisterProjectRoutes(router *gin.RouterGroup){
	projectRoutes := router.Group("/projects")
	
	projectRoutes.GET("/",controllers.GetProjects)
	projectRoutes.POST("/",controllers.InsertProject)
	projectRoutes.PATCH("/:id",controllers.UpdateProject)
	projectRoutes.DELETE("/:id",controllers.DeleteProject)

}