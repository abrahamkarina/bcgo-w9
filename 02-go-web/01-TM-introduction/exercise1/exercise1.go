package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]any{
			"message": "pong",
		})
	})

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
