package utils

import (
	controllers "github.com/ArshpreetS/Admin/handlers"
	"github.com/gin-gonic/gin"
)

func GetRoutes() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.GET("/test", func(c *gin.Context) {
			c.String(200, "It is working")
		})

		v1.POST("/add-bus", func(c *gin.Context) {
			controllers.HandlerAddBus(c)
		})

		v1.POST("/add-trip", func(c *gin.Context) {
			controllers.HandlerAddTrip(c)
		})

		v1.GET("/list-buses", func(c *gin.Context) {
			controllers.ListBuses(c)
		})
	}
	return r
}
