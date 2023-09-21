package app

import (
	"github.com/abrahamkarina/bcgo-w9/02-go-web/01-TT/exercise1n2/cmd/controllers"
	"github.com/abrahamkarina/bcgo-w9/02-go-web/01-TT/exercise1n2/internal/logger"
	"github.com/gin-gonic/gin"
)

type App struct {
	controllerProduct  *controllers.ProductController
	controllerPingPong *controllers.ControllerPingPong
	rt                 *gin.Engine
	log                logger.Logger
}

func NewApp(controllerProduct *controllers.ProductController, log logger.Logger, controllerPingPong *controllers.ControllerPingPong,
	ctx *gin.Engine) *App {

	app := new(App)
	app.log = log
	app.rt = ctx
	app.controllerProduct = controllerProduct
	app.controllerPingPong = controllerPingPong
	app.BuildRoutes()
	return app
}

func (app *App) BuildRoutes() {
	app.rt.Use(gin.Recovery())
	app.rt.Use(func(ctx *gin.Context) {
		app.log.Log(logger.Info{
			Method: ctx.Request.Method,
			Url:    ctx.Request.URL.Path,
		})
		ctx.Next()

	})

	app.controllerProduct.BuildRoutes(app.rt)
	app.controllerPingPong.BuildRoutes(app.rt)
}

func (app *App) Run() error {
	if err := app.rt.Run(":8080"); err != nil {
		return err
	}
	return nil
}
