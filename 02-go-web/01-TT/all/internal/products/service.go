package products

import (
	"errors"
	"github.com/abrahamkarina/bcgo-w9/02-go-web/01-TT/exercise1n2/internal/domain/model"
	"sort"
)

var (
	ErrIDInUse = errors.New("the id is already in use")
)

type ProductService struct {
	productRepo *ProductRepository
}

func NewProductService(productRepo *ProductRepository) *ProductService {
	svc := &ProductService{productRepo: productRepo}
	return svc
}

func (svc *ProductService) SearchByPrice(priceGt float64) ([]model.Product, error) {
	return svc.productRepo.SearchByPrice(priceGt)

}

func (svc *ProductService) GetById(id int) (*model.Product, error) {
	return svc.productRepo.GetById(id)

}

func (svc *ProductService) GetAll() ([]model.Product, error) {
	return svc.productRepo.GetAll()
}

func (svc *ProductService) isIDUsed(id int) (bool, error) {
	product, err := svc.GetById(id)
	if err != nil {
		return false, nil
	}
	if product != nil {
		return true, nil
	}
	return false, nil
}

func (svc *ProductService) Save(product model.Product) error {
	if product.ID == 0 {
		products, err := svc.productRepo.GetAll()
		if err != nil {
			return err
		}
		sortProductsByID(products)
		product.ID = products[len(products)-1].ID + 1
		return svc.productRepo.Save(product)
	}
	isUsed, err := svc.isIDUsed(product.ID)
	if err != nil {
		return err
	}
	if isUsed {
		return ErrIDInUse
	}
	return svc.productRepo.Save(product)
}

func (svc *ProductService) Put(product model.Product) error {
	return svc.productRepo.Put(product)
}

func (svc *ProductService) Delete(id int) error {
	return svc.productRepo.Delete(id)
}

func sortProductsByID(products []model.Product) {
	sort.Slice(products, func(i, j int) bool {
		return products[i].ID < products[j].ID
	})
}
