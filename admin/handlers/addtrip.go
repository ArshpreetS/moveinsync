package controllers

import (
	"context"
	"log"
	"net/http"

	db "github.com/ArshpreetS/Admin/DB"
	"github.com/ArshpreetS/Admin/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func HandlerAddTrip(c *gin.Context) {
	db_client := db.GetDBClient()
	defer db_client.Disconnect(context.Background())

	data := models.Bus_Trip{
		TripID: uuid.New().String(),
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, struct {
			Error  string `json:"error"`
			Status int    `json:"status"`
		}{
			Error:  err.Error(),
			Status: 0,
		})
	}

	// Checking if the bus exists

	all_buses_col := db_client.Database("Bus").Collection("all_buses")
	filter := bson.M{"busid": data.BusID}

	cursor, err := all_buses_col.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
		c.String(http.StatusInternalServerError, "Some error while searching in all_buses")
		return
	}
	defer cursor.Close(context.Background())

	var all_buses []models.Bus_data
	for cursor.Next(context.Background()) {
		var b models.Bus_data
		if err := cursor.Decode(&b); err != nil {
			log.Fatal(err)
		}
		all_buses = append(all_buses, b)
	}

	if len(all_buses) == 0 {
		c.String(http.StatusBadRequest, "No such bus exists already")
		return
	}

	if len(data.Route) < 2 || data.Route[0].DateTime != data.StartTime || data.Route[len(data.Route)-1].DateTime != data.EndTime {
		c.String(http.StatusBadRequest, "Route is not right")
		return
	}
	// Bus is there

	trip_col := db_client.Database("Bus").Collection("trips")

	// Adding the trip in Collection
	if _, err := trip_col.InsertOne(context.TODO(), data); err != nil {
		c.JSON(http.StatusInternalServerError, struct {
			Error  string `json:"error"`
			Status int    `json:"status"`
		}{
			Error:  err.Error(),
			Status: 0,
		})
	}

	c.JSON(http.StatusOK, struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		Status:  1,
		Message: "Trip added",
	})
}
