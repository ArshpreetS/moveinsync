package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ArshpreetS/moveinsync/gateway/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	godotenv.Load("./.env")

	{
		router.GET("/api/v1/test", func(c *gin.Context) {
			c.String(http.StatusOK, "Test is working")
		})

		router.POST("/api/v1/create-user", func(c *gin.Context) {
			handlers.CreateUser(c)
		})

		router.GET("/api/v1/:servicename/:service", func(c *gin.Context) {
			serviceName := c.Param("servicename")
			service := c.Param("service")

			switch serviceName {
			case "admin":
				res, err := http.Get(fmt.Sprint(os.Getenv("ADMINADDR"), ":", os.Getenv("ADMINPORT"), "/api/v1/", service))
				if err != nil {
					fmt.Println(err)
					c.String(http.StatusInternalServerError, "Error sending request to admin")
					return
				}
				defer res.Body.Close()

				// Read the response body from the GET request
				responseBody, err := io.ReadAll(res.Body)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
					return
				}

				// Forward the response from the GET request as the output of the Gin handler
				c.Data(res.StatusCode, res.Header.Get("Content-Type"), responseBody)
			case "user":

				res, err := http.Get(fmt.Sprint(os.Getenv("USERADDR"), ":", os.Getenv("ADMINPORT"), "/api/v1/", service))
				if err != nil {
					fmt.Println(err)
					c.String(http.StatusInternalServerError, "Error sending request to admin")
					return
				}
				defer res.Body.Close()

				// Read the response body from the GET request
				responseBody, err := io.ReadAll(res.Body)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
					return
				}

				// Forward the response from the GET request as the output of the Gin handler
				c.Data(res.StatusCode, res.Header.Get("Content-Type"), responseBody)
			}
		})

		router.POST("/api/v1/:servicename/:service", func(c *gin.Context) {
			serviceName := c.Param("servicename")
			service := c.Param("service")

			byteData, err := io.ReadAll(c.Request.Body)
			if err != nil {
				fmt.Println(err)
				c.String(http.StatusBadRequest, "Cannot read the body")
				return
			}
			switch serviceName {
			case "admin":

				res, err := http.Post(fmt.Sprint(os.Getenv("ADMINADDR"), ":", os.Getenv("ADMINPORT"), "/api/v1/", service), "application/json", bytes.NewBuffer(byteData))
				if err != nil {
					fmt.Println(err)
					c.String(http.StatusInternalServerError, "Error sending request to admin")
					return
				}
				defer res.Body.Close()

				responseData, err := io.ReadAll(res.Body)
				if err != nil {
					fmt.Println(err)
					c.String(http.StatusInternalServerError, "Error reading the output")
					return
				}
				c.Data(res.StatusCode, res.Header.Get("Content-Type"), responseData)
			case "user":

				res, err := http.Post(os.Getenv("USERADDR")+":"+os.Getenv("USERPORT")+"/api/v1/"+string(service), "application/json", bytes.NewBuffer(byteData))
				if err != nil {
					fmt.Println(err)
					c.String(http.StatusInternalServerError, "Error sending request to user")
					return
				}
				defer res.Body.Close()

				responseData, err := io.ReadAll(res.Body)
				if err != nil {
					fmt.Println(err)
					c.String(http.StatusInternalServerError, "Error reading the output")
					return
				}
				c.Data(res.StatusCode, res.Header.Get("Content-Type"), responseData)
			}

		})

	}
	router.Run(os.Getenv("ADDR"))
}
