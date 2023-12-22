package models

type Form struct {
	Name        string `json:"name" bson:"name"`
	Email       string `json:"email" bson:"email"`
	Description string `json:"description" bson:"description"`
	Timestamp string `json:"timestamp" bson:"timestamp"`
}
