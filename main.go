package main

import (
	

	"github.com/nishanth-code/portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

func init(){
	
	utils.ConnectToDB()


}

func test(c *gin.Context){
	c.JSON(200,gin.H{"message":"hello working fine"})
}

func main(){
r:=gin.Default()
r.GET("/",test)
r.Run()
}