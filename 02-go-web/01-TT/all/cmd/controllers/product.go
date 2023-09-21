package controllers

import (
	"github.com/abrahamkarina/bcgo-w9/02-go-web/01-TT/exercise1n2/cmd/config"
	"github.com/abrahamkarina/bcgo-w9/02-go-web/01-TT/exercise1n2/cmd/dto"
	"github.com/abrahamkarina/bcgo-w9/02-go-web/01-TT/exercise1n2/cmd/web"
	productsP "github.com/abrahamkarina/bcgo-w9/02-go-web/01-TT/exercise1n2/internal/products"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type ProductController struct {
	svc *productsP.ProductService
	cfg config.ProductConfig
}

func NewProductController(svc *productsP.ProductService, cfg config.ProductConfig) *ProductController {
	p := new(ProductController)
	p.cfg = cfg
	p.svc = svc
	return p
}

func (p *ProductController) save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request, ok := ctx.Get("dto")
		if !ok {
			ctx.JSON(http.StatusInternalServerError, web.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: "Error getting request data",
			})
			ctx.Abort()
			return
		}
		product := request.(dto.NewProduct).ToModel()
		err := p.svc.Save(product)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: "Error saving data",
			})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, web.Response{Data: "Ok"})
	}
}

func (p *ProductController) delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const idKey = "id"
		id, err := ParseToInt(ctx.Param(idKey))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{
				Status:  http.StatusBadRequest,
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: "Wrong param",
			})
			return
		}
		err = p.svc.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: "Error deleting",
			})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, web.Response{Data: "Ok"})
	}
}
func (p *ProductController) put() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request, ok := ctx.Get("dto")
		if !ok {
			ctx.JSON(http.StatusInternalServerError, web.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: "Error getting request data",
			})
			ctx.Abort()
			return
		}
		product := request.(dto.PutProduct).ToModel()
		err := p.svc.Put(product)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: "Error in put",
			})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, web.Response{Data: "Ok"})
	}
}

func (p *ProductController) validateAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		ok := p.validateToken(token)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, web.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Code:    strconv.Itoa(http.StatusUnauthorized),
				Message: "Invalid token",
			})
			ctx.Abort()
			return

		}
		ctx.Next()
	}
}

func (p *ProductController) validateContractNew() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var requestBody dto.NewProduct
		err := ctx.ShouldBindJSON(&requestBody)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{
				Status:  http.StatusBadRequest,
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: "Invalid request body",
			})
			ctx.Abort()
			return
		}

		ok := isDateValid(requestBody.Expiration)
		if !ok {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{
				Status:  http.StatusBadRequest,
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: "Invalid date format",
			})
			ctx.Abort()
			return
		}
		ctx.Set("dto", requestBody)
		ctx.Next()
	}
}
func (p *ProductController) validateContractUpdate() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var requestBody dto.PatchProduct
		err := ctx.ShouldBindJSON(&requestBody)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{
				Status:  http.StatusBadRequest,
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: "Invalid request body",
			})
			ctx.Abort()
			return
		}

		ok := isDateValid(*requestBody.Expiration)
		if !ok {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{
				Status:  http.StatusBadRequest,
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: "Invalid date format",
			})
			ctx.Abort()
			return
		}
		ctx.Set("dto", requestBody)
		ctx.Next()
	}
}

func (p *ProductController) validateToken(token string) bool {
	if token == p.cfg.Token {
		return true
	}
	return false
}

func (p *ProductController) getAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := p.svc.GetAll()
		if err != nil {

		}
		ctx.JSON(http.StatusOK, products)
		return
	}
}

func (p *ProductController) getByID() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		const idKey = "id"
		id, err := ParseToInt(ctx.Param(idKey))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{
				Status:  http.StatusBadRequest,
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: "Wrong param",
			})
			return
		}
		product, err := p.svc.GetById(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: "Error searching by id",
			})
			return
		}
		ctx.JSON(http.StatusOK, product)
		return
	}
}

func (p *ProductController) search() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		const priceKey = "priceGt"
		price, err := ParseToFloat(ctx.Query(priceKey))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{
				Status:  http.StatusBadRequest,
				Code:    strconv.Itoa(http.StatusBadRequest),
				Message: "Invalid param",
			})
			return
		}
		product, err := p.svc.SearchByPrice(price)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Code:    strconv.Itoa(http.StatusInternalServerError),
				Message: "Internal error searching by price",
			})
			return
		}
		ctx.JSON(http.StatusOK, product)
		return
	}
}

func (p *ProductController) BuildRoutes(rt *gin.Engine) {
	pGr := rt.Group("/products")
	pGr.Use(p.validateAuth())
	pGr.POST("", p.validateContractNew(), p.save())
	//pGr.PUT("/:id", p.validateContractUpdate(), p.put())
	//pGr.PATCH("/:id", p.validateContractUpdate(), p.patch())
	pGr.GET("", p.getAll())
	pGr.GET("/:id", p.getByID())
	pGr.DELETE("/:id", p.validateAuth(), p.delete())
	pGr.GET("/search", p.search())
}

func ParseToFloat(toConvert string) (float64, error) {
	return strconv.ParseFloat(toConvert, 64)
}

func ParseToInt(toConvert string) (int, error) {
	return strconv.Atoi(toConvert)
}

func isDateValid(dateStr string) bool {
	layout := "02/01/2006"
	_, err := time.Parse(layout, dateStr)
	return err == nil
}
