package main

import (
	"github.com/nishanth-code/portfolio-backend/auth"
	"github.com/nishanth-code/portfolio-backend/routes"
	"github.com/nishanth-code/portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

func init(){
	utils.LoadEnvVariables()
	utils.ConnectToDB()
}


func main(){
router:=gin.Default()
router.POST("/admin/auth",auth.Authenticate)

api:=router.Group("/api")
{
	routes.RegisterProjectRoutes(api)
	routes.RegisterExperienceRoutes(api)
	routes.RegisterTestimonialRoutes(api)
}
router.Run()
}