package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(host, port, db string) (*mongo.Database, error) {
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port))
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	return client.Database(db), nil
}
