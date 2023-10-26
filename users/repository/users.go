package repository

import (
	"context"
	"fmt"
	"users/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRepository(db mongo.Database) Repository {
	return Repository{DB: &db}
}

type Repository struct {
	DB *mongo.Database
}

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

func (r *Repository) FindUserbyEmail(email string) (entity.Users, error) {
	collection := r.DB.Collection("users")

	filter := bson.M{"email": email}

	var user entity.Users
	if err := collection.FindOne(context.Background(), filter).Decode(&user); err != nil {
		return entity.Users{}, fmt.Errorf(err.Error())
	}

	return user, nil
}

func (r *Repository) FindUserById(id primitive.ObjectID) (entity.Users, error) {
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
			"username":     reqbodyUser.Username,
			"email":        reqbodyUser.Email,
			"phone_number": reqbodyUser.Phone_number,
			"age":          reqbodyUser.Age,
			"description":  reqbodyUser.Description,
			"image_url":    reqbodyUser.Image_url,
		},
	}

	_, err := collection.UpdateOne(context.Background(), filter, updateduserDoc)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}
