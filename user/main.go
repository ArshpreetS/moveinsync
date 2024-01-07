package main

import (
	"fmt"
	"net/http"
	"os"

	controllers "github.com/ArshpreetS/moveinsync/user/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getRoutes() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	{
		v1.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, struct {
				Message string `json:"message"`
			}{
				Message: "Test is working",
			})
		})

		v1.POST("/get-buses", func(c *gin.Context) {
			controllers.GetBuses(c)
		})

		v1.POST("/book-ticket", func(c *gin.Context) {
			controllers.BookTickets(c)
		})
	}

	return router
}

func main() {
	router := getRoutes()

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading the environement file")
	}

	router.Run(os.Getenv("ADDR"))
}
