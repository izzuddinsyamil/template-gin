package repo

import "go.mongodb.org/mongo-driver/mongo"

type mongoRepo struct {
	db *mongo.Database
}

func NewMongoRepo(db *mongo.Database) *mongoRepo {
	return &mongoRepo{
		db: db,
	}
}
