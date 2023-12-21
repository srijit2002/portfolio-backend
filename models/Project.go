package models

type Project struct {
	Name        string             `json:"name" bson:"name"`
	Url         string             `json:"url" bson:"url"`
	Description string             `json:"description" bson:"description"`
	Image       string             `json:"image" bson:"image"`
}
