package utils

import(
	"context"
	"log"
	"fmt"
	
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
)
func ConnectToDB(){
	uri:= "mongodb+srv://nishanth:nish9741@cluster0.bbodeek.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	clientOptions := options.Client().ApplyURI(uri)
	Client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer Client.Disconnect(context.TODO())

	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	
}

