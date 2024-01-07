package controllers

import (
	"context"
	"fmt"
	"net/http"

	db "github.com/ArshpreetS/Admin/DB"
	"github.com/ArshpreetS/Admin/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ListBuses(c *gin.Context) {
	db_client := db.GetDBClient()
	defer db_client.Disconnect(context.Background())

	col := db_client.Database("Bus").Collection("all_buses")

	cursor, err := col.Find(context.TODO(), bson.D{})

	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "Cannot find documents in buses collection")
		return
	}

	var results []models.Bus_data
	for cursor.Next(context.Background()) {
		var data models.Bus_data

		if err := cursor.Decode(&data); err != nil {
			fmt.Println(err)
			c.String(http.StatusInternalServerError, "Error while decoding document")
			return
		}

		results = append(results, data)
	}

	c.JSON(http.StatusOK, struct {
		AllBuses []models.Bus_data `json:"allBuses"`
		Status   int               `json:"status"`
	}{
		AllBuses: results,
		Status:   1,
	})
}
