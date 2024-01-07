package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ArshpreetS/moveinsync/user/database"
	"github.com/ArshpreetS/moveinsync/user/models"
	"github.com/ArshpreetS/moveinsync/user/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetBuses(c *gin.Context) {
	db_client := database.GetDBClient()
	defer db_client.Disconnect(context.Background())

	travelData := models.TravelData{}

	if err := c.ShouldBindJSON(&travelData); err != nil {
		fmt.Println(err.Error() + "Error while binding")
		c.AbortWithStatusJSON(http.StatusBadRequest, struct {
			Error  string `json:"error"`
			Status int    `json:"status"`
		}{
			Error:  err.Error(),
			Status: 0,
		})
		return
	}

	trips_col := db_client.Database("Bus").Collection("trips")
	bus_col := db_client.Database("Bus").Collection("all_buses")
	tickets_col := db_client.Database("Bus").Collection("tickets")

	// getting all trips
	cursor, err := trips_col.Find(context.Background(), bson.D{})
	if err != nil {
		fmt.Println("Error finding documents in mongodb")
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, struct {
			Error  string `json:"error"`
			Status int    `json:"status"`
		}{
			Error:  err.Error(),
			Status: 0,
		})
	}

	defer cursor.Close(context.Background())

	var results []models.Response

	for cursor.Next(context.Background()) {
		var data models.Response

		if err := cursor.Decode(&data); err != nil {
			log.Fatal("Can't decode the data from trips")
		}

		toAdd := false
		foundSource := false
		var start_time time.Time
		var end_time time.Time
		for _, val := range data.Route {
			if val.Stop == travelData.Source &&
				time.Date(val.DateTime.Year(), val.DateTime.Month(), val.DateTime.Day(), 0, 0, 0, 0, val.DateTime.Location()) ==
					time.Date(travelData.Date.Year(), travelData.Date.Month(), travelData.Date.Day(), 0, 0, 0, 0, travelData.Date.Location()) {
				foundSource = true
				start_time = val.DateTime
			}
			if foundSource && val.Stop == travelData.Destination {
				toAdd = true
				end_time = val.DateTime
			}
		}
		if toAdd {

			// Finding all seats
			data.Tickets, err = utils.GetSeats(data.BusID, bus_col)
			if err != nil {
				fmt.Println(err)
				c.String(http.StatusInternalServerError, "There is error while binding document")
				return
			}

			// Removing the used tickets
			filter := bson.M{"tripid": data.TripID}

			cur, err := tickets_col.Find(context.TODO(), filter)
			if err != nil {
				fmt.Println(err)
				c.String(http.StatusInternalServerError, "Couldn't find tickets documents")
				return
			}

			for cur.Next(context.Background()) {
				var ticket models.Ticket
				if err := cur.Decode(&ticket); err != nil {
					fmt.Println(err)
					c.String(http.StatusInternalServerError, "Couldn't bind the ticket")
					return
				}

				if ticket.StartTime.Before(end_time) && start_time.Before(ticket.EndTime) {
					data.Tickets--
				}
			}

			results = append(results, models.Response{
				TripID:    data.TripID,
				BusID:     data.BusID,
				Route:     data.Route,
				Tickets:   data.Tickets,
				StartTime: start_time,
				EndTime:   end_time,
			})
		}
	}

	// remove all tickets having 0 tickets
	results = utils.FilterZeroTickets(results)

	// results consist of possible trips
	c.JSON(http.StatusOK, struct {
		All_buses []models.Response
	}{
		All_buses: results,
	})

}
