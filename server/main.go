package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	middlewares "github.com/tomi-saku/jsys25-advent-calender/middleware"
	"github.com/tomi-saku/jsys25-advent-calender/models"
)

func main() {
	fmt.Println("Application Initializing...")

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{
			"GET",
			"POST",
			"OPTIONS",
			"DELETE",
		},
		AllowHeaders: []string{
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/health", getHealth)

	authorizedRoutes := r.Group("/authorization")
	authorizedRoutes.Use(middlewares.AuthMiddleware(os.Getenv("GOOGLE_CLIENT_ID")))
	{
		authorizedRoutes.GET("", getEmailAddress)
	}

	fmt.Println("Application Starts!")
	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080
}

func getHealth(c *gin.Context) {
	message := models.Message{
		Message: "Hello, World!",
	}
	c.IndentedJSON(http.StatusOK, message)
}

func getEmailAddress(c *gin.Context) {
	email := c.GetString("email")
	image := c.GetString("image")
	res := models.User{
		Email: email,
		Image: image,
	}
	c.IndentedJSON(http.StatusOK, res)
}
