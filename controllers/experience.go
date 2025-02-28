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

func GetPExperience(c *gin.Context){
	var result  []models.Experience
	ctx ,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	err := utils.ExperienceCollection.Find(ctx,bson.M{},&result)
	if err!=nil{
		c.JSON(500,gin.H{"error": "Failed to fetch experience"})
	}
	c.JSON(200,gin.H{"data":result})

}

func InsertExperience(c *gin.Context){
	var reqBody struct{
		Year  int                 `json:"year"`
		Role   string                 `json:"role"`
		Organization  string                 `json:"organization"`
		Description string                 `json:"description"`
		
	} 
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON request body"})
		return
	}

	if reqBody.Year == 0 || reqBody.Role == "" || reqBody.Organization == "" || 
		reqBody.Description == "" {
		c.JSON(400, gin.H{"error": "All fields are required"})
		return
	}
	project := models.Experience{
		Year: reqBody.Year,
		Role: reqBody.Role,
		Organization: reqBody.Organization,
		Description: reqBody.Description,
		}
	ctx ,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	result,err := utils.ExperienceCollection.InsertOne(ctx,project)
	if err!=nil{
		c.JSON(500,gin.H{"error": "Failed to insert experience"})
	}

	c.JSON(200,gin.H{"data":result})

}

func UpdateExperience(c *gin.Context){
	var reqBody struct{
		Year  int                 `json:"year"`
		Role   string                 `json:"role"`
		Organization  string                 `json:"organization"`
		Description string                 `json:"description"`
		
	} 

	projectID := c.Param("id")
	if projectID == "" {
		c.JSON(400, gin.H{"error": "experience ID is required"})
		return
	}

	objID, err := primitive.ObjectIDFromHex(projectID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid experience ID"})
		return
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON request body"})
		return
	}

	if reqBody.Year == 0 || reqBody.Role == "" || reqBody.Organization == "" || 
		reqBody.Description == "" {
		c.JSON(400, gin.H{"error": "All fields are required"})
		return
	}

	update := bson.M{
		"$set": bson.M{
		"year": reqBody.Year,
		"role": reqBody.Role,
		"organization": reqBody.Organization,
		"description": reqBody.Description,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	
	result, err := utils.ExperienceCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update experience"})
		return
	}

	
	if result.MatchedCount == 0 {
		c.JSON(404, gin.H{"error": "experience not found"})
		return
	}

	// Success response
	c.JSON(200, gin.H{"message": "experience updated successfully"})
}

func DeleteExperience(c *gin.Context){
	experienceId := c.Param("id")

	if experienceId == "" {
		c.JSON(400, gin.H{"error": "experience ID is required"})
		return
	}

	objID, err := primitive.ObjectIDFromHex(experienceId)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid experience ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result,err := utils.ExperienceCollection.DeleteOne(ctx,bson.M{"_id":objID})
	if err!=nil{
		c.JSON(500,gin.H{"error":"failed to delete experience"})
		return
	}

	if result.DeletedCount == 0{
		c.JSON(404, gin.H{"error": "experience not found"})
		return
	
	}

	c.JSON(200, gin.H{"message": "experience deleted successfully"})

}
