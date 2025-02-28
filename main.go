package main

import (
	

	"github.com/nishanth-code/portfolio-backend/utils"
	"github.com/nishanth-code/portfolio-backend/routes"

	"github.com/gin-gonic/gin"
)

func init(){
	
	utils.ConnectToDB()
}

// func test(c *gin.Context){
// 	c.JSON(200,gin.H{"message":"hello working fine"})
// }

func main(){
router:=gin.Default()
// projects:=router.Group("/projects")
api:=router.Group("/api")
{
	routes.RegisterProjectRoutes(api)
	routes.RegisterExperienceRoutes(api)
}
router.Run(":3000")
}