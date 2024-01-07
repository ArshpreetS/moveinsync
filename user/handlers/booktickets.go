package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ArshpreetS/moveinsync/user/database"
	"github.com/ArshpreetS/moveinsync/user/models"
	"github.com/gin-gonic/gin"
)

func BookTickets(c *gin.Context) {
	db_client := database.GetDBClient()
	defer db_client.Disconnect(context.Background())

	var data models.TicketToBook

	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "There was some error with binding body with json")
		return
	}

	ticket_col := db_client.Database("Bus").Collection("tickets")

	_, err := ticket_col.InsertOne(context.Background(), data)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "There was error while adding document")
		return
	}

	c.JSON(http.StatusAccepted, struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		Status:  1,
		Message: "Tickets Booked",
	})
}
