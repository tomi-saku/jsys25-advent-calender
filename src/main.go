package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	middlewares "github.com/tomi-saku/jsys25-advent-calender/middleware"
	"github.com/tomi-saku/jsys25-advent-calender/models"
)

func main() {
	fmt.Println("Application Initializing...")

	r := gin.Default()

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
	userId := c.GetString("userID")
	res := models.UserId{
		UserId: userId,
	}
	c.IndentedJSON(http.StatusOK, res)
}
