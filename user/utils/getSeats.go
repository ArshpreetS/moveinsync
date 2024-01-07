package utils

import (
	"context"

	"github.com/ArshpreetS/moveinsync/user/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSeats(busid string, db *mongo.Collection) (int, error) {
	filter := bson.M{"busid": busid}
	cur := db.FindOne(context.TODO(), filter)
	bus_d := models.Bus_data{}
	if err := cur.Decode(&bus_d); err != nil {
		return 0, err
	}
	return bus_d.Seats, nil
}
