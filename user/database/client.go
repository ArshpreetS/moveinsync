package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDBClient() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://"+os.Getenv("DBUSER")+":"+os.Getenv("DBPASS")+"@localhost:27017"))
	if err != nil {
		fmt.Println("Issue connecting to mongodb")
		panic(err)
	}
	return client
}
