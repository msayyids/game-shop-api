package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	DB *mongo.Database
}

func NewRepository(db mongo.Database) Repository {
	return Repository{DB: &db}
}
