package repo

import (
	"context"
	"template-gin/model"

	"go.mongodb.org/mongo-driver/bson"
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

func (r *mongoRepo) GetBooks(ctx context.Context, title string, startYear, endYear int) (res []model.Book, err error) {
	filter := bson.M{}

	if title != "" {
		filter["title"] = title
	}

	if startYear != 0 {
		filter["release_year"] = bson.M{
			"$gte": startYear,
		}
	}

	if endYear != 0 {
		filter["release_year"] = bson.M{
			"$lt": endYear,
		}
	}

	cursor, err := r.db.Collection("books").Find(ctx, filter)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var row model.Book
		err = cursor.Decode(&row)
		if err != nil {
			return
		}

		res = append(res, row)
	}

	return
}
