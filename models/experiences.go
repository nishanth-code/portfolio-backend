package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	
)

type Experience struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Year int `bson:"year"`
	Role string `bson:"role"`
	Organization string `bson:"organization"`
	Description string `bson:"description"`
}