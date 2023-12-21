package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Url         string             `json:"url" bson:"url"`
	Description string             `json:"description" bson:"description"`
	Image       []byte             `json:"image" bson:"image"`
}
