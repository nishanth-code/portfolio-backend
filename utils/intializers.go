package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client     *mongo.Client
	ProjectCollection    *Collection
	ExperienceCollection *Collection
	TestimonialCollection *Collection
)
func ConnectToDB(){
	uri:= os.Getenv("DBUri")

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	db := client.Database(os.Getenv("DBName"))

	Client = client
	ProjectCollection = &Collection{db.Collection("projects")}
	ExperienceCollection = &Collection{db.Collection("experience")}
	TestimonialCollection = &Collection{db.Collection("testimonial")}
	
}

func Close() {
	if Client != nil {
		err := Client.Disconnect(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Disconnected from MongoDB!")
	}
}

func LoadEnvVariables(){
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

}

