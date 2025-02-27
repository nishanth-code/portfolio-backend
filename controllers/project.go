package controllers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nishanth-code/portfolio-backend/models"
	"github.com/nishanth-code/portfolio-backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func GetProjects(c *gin.Context){
	var result  []models.Project
	ctx ,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	err := utils.ProjectCollection.Find(ctx,bson.M{},&result)
	if err!=nil{
		c.JSON(500,gin.H{"error": "Failed to insert project"})
	}
	c.JSON(200,gin.H{"data":result})

}

func InsertProject(c *gin.Context){
	var reqBody struct{
		DisplayPic  string                 `json:"displaypic"`
		PublicUrl   string                 `json:"publicUrl"`
		GithubUrl   string                 `json:"githubUrl"`
		ProjectName string                 `json:"projectName"`
		Description string                 `json:"description"`
		Domain      models.Specialization  `json:"domain"`

	} 
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON request body"})
		return
	}

	if reqBody.DisplayPic == "" || reqBody.PublicUrl == "" || reqBody.GithubUrl == "" || 
		reqBody.ProjectName == "" || reqBody.Description == "" || reqBody.Domain == "" {
		c.JSON(400, gin.H{"error": "All fields are required"})
		return
	}
	project := models.Project{
		DisplayPic: reqBody.DisplayPic,
		PublicUrl: reqBody.PublicUrl,
		GithubUrl: reqBody.GithubUrl,
		ProjectName: reqBody.ProjectName,
		Description: reqBody.Description,
		Domain: reqBody.Domain,
	}
	ctx ,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	result,err := utils.ProjectCollection.InsertOne(ctx,project)
	if err!=nil{
		c.JSON(500,gin.H{"error": "Failed to insert project"})
	}

	c.JSON(200,gin.H{"data":result})

}

func UpdateProject(c *gin.Context){
	var reqBody struct{
		DisplayPic  string                 `json:"displaypic"`
		PublicUrl   string                 `json:"publicUrl"`
		GithubUrl   string                 `json:"githubUrl"`
		ProjectName string                 `json:"projectName"`
		Description string                 `json:"description"`
		Domain      models.Specialization  `json:"domain"`

	} 

	projectID := c.Param("id")
	if projectID == "" {
		c.JSON(400, gin.H{"error": "Project ID is required"})
		return
	}

	objID, err := primitive.ObjectIDFromHex(projectID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid project ID"})
		return
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON request body"})
		return
	}

	if reqBody.DisplayPic == "" || reqBody.PublicUrl == "" || reqBody.GithubUrl == "" || 
		reqBody.ProjectName == "" || reqBody.Description == "" || reqBody.Domain == "" {
		c.JSON(400, gin.H{"error": "All fields are required"})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"displaypic":  reqBody.DisplayPic,
			"publicUrl":   reqBody.PublicUrl,
			"githubUrl":   reqBody.GithubUrl,
			"projectName": reqBody.ProjectName,
			"description": reqBody.Description,
			"domain":      reqBody.Domain,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	
	result, err := utils.ProjectCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update project"})
		return
	}

	
	if result.MatchedCount == 0 {
		c.JSON(404, gin.H{"error": "Project not found"})
		return
	}

	// Success response
	c.JSON(200, gin.H{"message": "Project updated successfully"})
}

func DeleteProject(c *gin.Context){
	projectId := c.Param("id")

	if projectId == "" {
		c.JSON(400, gin.H{"error": "Project ID is required"})
		return
	}

	objID, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid project ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result,err := utils.ProjectCollection.DeleteOne(ctx,bson.M{"_id":objID})
	if err!=nil{
		c.JSON(500,gin.H{"error":"failed to delete project"})
		return
	}

	if result.DeletedCount == 0{
		c.JSON(404, gin.H{"error": "Project not found"})
		return
	
	}

	c.JSON(200, gin.H{"message": "Project deleted successfully"})

}
