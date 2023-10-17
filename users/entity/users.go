package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Username     string             `json:"username" bson:"username"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password"`
	Phone_number string             `json:"phone_number" bson:"phone_number"`
	Age          string             `json:"age" bson:"age"`
	Description  string             `json:"description" bson:"description"`
	Image_url    string             `json:"image_url" bson:"imagae_url"`
}
