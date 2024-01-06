package controllers

import (
	"context"
	"net/http"

	db "github.com/ArshpreetS/Admin/DB"
	"github.com/ArshpreetS/Admin/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandlerAddBus(c *gin.Context) {
	db_client := db.GetDBClient()
	defer db_client.Disconnect(context.TODO())

	data := models.Bus_data{
		BusID: uuid.New().String(),
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

	col := db_client.Database("Bus").Collection("all_buses")

	if _, err := col.InsertOne(context.TODO(), data); err != nil {
		c.JSON(http.StatusInternalServerError, struct {
			Error   string `json:"error`
			Message string `json:"message"`
			Status  int    `json:"status"`
		}{
			Error:   err.Error(),
			Message: "Couldn't add data to mongodb",
			Status:  0,
		})
	}

	c.JSON(http.StatusOK, struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{
		Status:  1,
		Message: "Bus Added!",
	})
}
