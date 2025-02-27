package utils

import(
	"context"
	"log"
	"fmt"
	
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Client     *mongo.Client
	ProjectCollection    *Collection
	ExperienceCollection *Collection
	TestimonialCollection *Collection
)
func ConnectToDB(){
	uri:= "mongodb+srv://nishanth:nish9741@cluster0.bbodeek.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

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

	db := client.Database("portfolio")

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

