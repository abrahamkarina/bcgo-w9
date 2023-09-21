package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ControllerPingPong struct {
}

const message = "pong"

func NewControllerPingPong() *ControllerPingPong {
	return &ControllerPingPong{}
}
func (pingpong *ControllerPingPong) PingPong() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(http.StatusOK, message)
	}
}

func (pingpong *ControllerPingPong) BuildRoutes(rt *gin.Engine) {
	rt.GET("/pingpong", pingpong.PingPong())

}
