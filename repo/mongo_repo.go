package repo

import (
	"context"
	"template-gin/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepo struct {
	db *mongo.Database
}

func NewMongoRepo(db *mongo.Database) *mongoRepo {
	return &mongoRepo{
		db: db,
	}
}

func (r *mongoRepo) InsertBook(ctx context.Context, b model.Book) error {
	_, err := r.db.Collection("books").InsertOne(ctx, b)
	if err != nil {
		return err
	}

	return nil
}
