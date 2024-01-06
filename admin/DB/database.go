package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDBClient() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:root@localhost:27017"))
	if err != nil {
		panic(err)
	}
	return client
}
