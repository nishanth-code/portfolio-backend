package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Specialization string;

const(
	Web3  Specialization = "web3"
	Iot  Specialization = "iot"
	Webdevelopment Specialization = "webdevelopment"
)

type Project struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DisplayPic string `bson:"displaypic"`
	PublicUrl string `bson:"publicUrl"`
	GithubUrl string `bson:"githubUrl"`
	ProjectName string `bson:"projectName"`
	Description string `bson:"description"`
	Domain Specialization `bson:"domain"`
	

}