package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Testimonial struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PersonName string `bson:"personName"`
	Organization string `bson:"organization"`
	Description string `bson:"description"`
	ImageUrl string `bson:"imageUrl"`

}