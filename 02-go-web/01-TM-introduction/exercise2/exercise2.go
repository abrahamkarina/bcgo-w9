package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GreetingRequest struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

func main() {
	router := gin.Default()

	router.POST("/greeting", func(c *gin.Context) {
		var greetingRequest GreetingRequest

		if err := c.ShouldBindJSON(&greetingRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		name := greetingRequest.Name
		lastname := greetingRequest.Lastname
		respond := fmt.Sprintf("Hello %s %s", name, lastname)

		c.JSON(http.StatusOK, gin.H{"message": respond})
	})

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
