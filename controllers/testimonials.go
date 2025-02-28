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

func GetTestimonial(c *gin.Context){
	var result  []models.Testimonial
	ctx ,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	err := utils.TestimonialCollection.Find(ctx,bson.M{},&result)
	if err!=nil{
		c.JSON(500,gin.H{"error": "Failed to fetch testimonials"})
	}
	c.JSON(200,gin.H{"data":result})

}

func InsertTestimonial(c *gin.Context){
	var reqBody struct{
		PersonName   string `json:"personName"`
		Organization string `json:"organization"`
		Description  string `json:"description"`
		ImageUrl     string `json:"imageUrl"`
		
	} 
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON request body"})
		return
	}

	if reqBody.PersonName == "" || reqBody.ImageUrl == "" || reqBody.Organization == "" || 
		reqBody.Description == "" {
		c.JSON(400, gin.H{"error": "All fields are required"})
		return
	}
	project := models.Testimonial{
		PersonName: reqBody.PersonName,
		ImageUrl: reqBody.ImageUrl,
		Organization: reqBody.Organization,
		Description: reqBody.Description,
		}

	ctx ,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	result,err := utils.TestimonialCollection.InsertOne(ctx,project)
	if err!=nil{
		c.JSON(500,gin.H{"error": "Failed to insert experience"})
	}

	c.JSON(200,gin.H{"data":result})

}

func UpdateTestimonial(c *gin.Context){
	var reqBody struct{
		PersonName    string  `json:"personName"`
		Organization  string  `json:"organization"`
		Description   string  `json:"description"`
		ImageUrl      string  `json:"imageUrl"`
		
	} 

	testimonialID := c.Param("id")
	if testimonialID == "" {
		c.JSON(400, gin.H{"error": "testimonial ID is required"})
		return
	}

	objID, err := primitive.ObjectIDFromHex(testimonialID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid testimonial ID"})
		return
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON request body"})
		return
	}

	if reqBody.PersonName == "" || reqBody.ImageUrl == "" || reqBody.Organization == "" || 
	reqBody.Description == "" {
		c.JSON(400, gin.H{"error": "All fields are required"})
		return
	}

	update := bson.M{
		"$set": bson.M{
		"personName": reqBody.PersonName,
		"imageUrl": reqBody.ImageUrl,
		"organization": reqBody.Organization,
		"description": reqBody.Description,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	
	result, err := utils.TestimonialCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update testimonial"})
		return
	}

	
	if result.MatchedCount == 0 {
		c.JSON(404, gin.H{"error": "testimonial not found"})
		return
	}

	// Success response
	c.JSON(200, gin.H{"message": "testimonial updated successfully"})
}

func DeleteTestimonial(c *gin.Context){
	testimonialId := c.Param("id")

	if testimonialId == "" {
		c.JSON(400, gin.H{"error": "testimonial ID is required"})
		return
	}

	objID, err := primitive.ObjectIDFromHex(testimonialId)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid testimonial ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result,err := utils.TestimonialCollection.DeleteOne(ctx,bson.M{"_id":objID})
	if err!=nil{
		c.JSON(500,gin.H{"error":"failed to delete testimonial"})
		return
	}

	if result.DeletedCount == 0{
		c.JSON(404, gin.H{"error": "testimonial not found"})
		return
	
	}

	c.JSON(200, gin.H{"message": "testimonial deleted successfully"})

}
