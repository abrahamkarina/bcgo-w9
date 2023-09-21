package main

import (
	appP "github.com/abrahamkarina/bcgo-w9/02-go-web/01-TT/exercise1n2/cmd/app"
	"github.com/abrahamkarina/bcgo-w9/02-go-web/01-TT/exercise1n2/cmd/config"
	"github.com/abrahamkarina/bcgo-w9/02-go-web/01-TT/exercise1n2/cmd/controllers"
	"github.com/abrahamkarina/bcgo-w9/02-go-web/01-TT/exercise1n2/internal/logger"
	"github.com/abrahamkarina/bcgo-w9/02-go-web/01-TT/exercise1n2/internal/products"
	"github.com/abrahamkarina/bcgo-w9/02-go-web/01-TT/exercise1n2/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	ctx := gin.New()

	db := store.NewStore()
	err := db.Init("products.json")
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}
	productRepo, err := products.NewProductRepository(db)
	if err != nil {
		panic(err)
	}
	productSvc := products.NewProductService(productRepo)
	controllerProduct := controllers.NewProductController(productSvc, getProductCfg())
	controllerPingPong := controllers.NewControllerPingPong()

	app := appP.NewApp(controllerProduct, logger.NewLogger(), controllerPingPong, ctx)
	if err := app.Run(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

}

func getProductCfg() config.ProductConfig {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	token := os.Getenv("TOKEN")
	return config.ProductConfig{Token: token}

}
