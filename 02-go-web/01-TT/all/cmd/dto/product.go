package dto

import "github.com/abrahamkarina/bcgo-w9/02-go-web/01-TT/exercise1n2/internal/domain/model"

type Dto interface {
	ToModel() *model.Product
}

type NewProduct struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Expiration  string  `json:"expiration" binding:"required"`
	IsPublished bool    `json:"is_published"`
	Price       float64 `json:"price" binding:"required"`
	Quantity    int     `json:"quantity" `
}

func (dto NewProduct) ToModel() model.Product {
	product := model.Product{
		ID:          dto.ID,
		Name:        dto.Name,
		Quantity:    dto.Quantity,
		IsPublished: dto.IsPublished,
		Expiration:  dto.Expiration,
		Price:       dto.Price,
	}

	return product
}

type PatchProduct struct {
	ID          *int     `json:"id" binding:"required"`
	Name        *string  `json:"name" binding:"required"`
	Expiration  *string  `json:"expiration" binding:"required"`
	IsPublished *bool    `json:"is_published"`
	Price       *float64 `json:"price" binding:"required"`
	Quantity    *int     `json:"quantity" `
}

func (dto PatchProduct) ToModel() model.Product {
	product := model.Product{
		Name:        *dto.Name,
		Quantity:    *dto.Quantity,
		IsPublished: *dto.IsPublished,
		Expiration:  *dto.Expiration,
		Price:       *dto.Price,
	}

	return product
}

type PutProduct struct {
	ID          *int     `json:"id" binding:"required"`
	Name        *string  `json:"name" binding:"required"`
	Expiration  *string  `json:"expiration" binding:"required"`
	IsPublished *bool    `json:"is_published"`
	Price       *float64 `json:"price" binding:"required"`
	Quantity    *int     `json:"quantity" `
}

func (dto PutProduct) ToModel() model.Product {
	product := model.Product{
		Name:        *dto.Name,
		Quantity:    *dto.Quantity,
		IsPublished: *dto.IsPublished,
		Expiration:  *dto.Expiration,
		Price:       *dto.Price,
	}

	return product
}
