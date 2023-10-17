package repository

import (
	"context"
	"fmt"
	"users/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) Adduser(reqbodyUser entity.Users) (entity.Users, error) {
	collection := r.DB.Collection("users")

	newUserDoc := entity.Users{
		ID:           primitive.NewObjectID(),
		Username:     reqbodyUser.Username,
		Email:        reqbodyUser.Email,
		Password:     reqbodyUser.Password,
		Phone_number: reqbodyUser.Phone_number,
		Age:          reqbodyUser.Age,
		Description:  reqbodyUser.Description,
		Image_url:    reqbodyUser.Image_url,
	}

	_, err := collection.InsertOne(context.Background(), newUserDoc)
	if err != nil {
		return entity.Users{}, nil
	}

	return newUserDoc, nil
}

func (r *Repository) Getuser(id primitive.ObjectID) (entity.Users, error) {
	collection := r.DB.Collection("users")
	filter := bson.M{"_id": id}

	var user entity.Users
	if err := collection.FindOne(context.Background(), filter).Decode(&user); err != nil {
		return entity.Users{}, fmt.Errorf(err.Error())
	}

	return user, nil
}

func (r *Repository) EditUser(id primitive.ObjectID, reqbodyUser entity.Users) error {
	collection := r.DB.Collection("users")

	filter := bson.M{"_id": id}
	updateduserDoc := bson.M{
		"$set": bson.M{
			"phone_number": reqbodyUser.Phone_number,
			"age":          reqbodyUser.Age,
			"description":  reqbodyUser.Description,
			"image_url":    reqbodyUser.Image_url,
		},
	}

	if err := collection.FindOneAndUpdate(context.Background(), filter, updateduserDoc); err != nil {
		return err.Err()
	}

	return nil
}
