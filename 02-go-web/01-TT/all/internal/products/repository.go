package products

import (
	"github.com/abrahamkarina/bcgo-w9/02-go-web/01-TT/exercise1n2/internal/domain/model"
	"github.com/abrahamkarina/bcgo-w9/02-go-web/01-TT/exercise1n2/pkg/store"
)

type ProductRepository struct {
	db store.Store
}

func NewProductRepository(db store.Store) (*ProductRepository, error) {
	r := new(ProductRepository)
	r.db = db
	return r, nil
}

func (repo *ProductRepository) SearchByPrice(priceGt float64) ([]model.Product, error) {
	products, err := repo.db.GetAll()
	if err != nil {
		return nil, err
	}
	respond := make([]model.Product, 0)
	for _, product := range products {
		if product.Price > priceGt {
			respond = append(respond, product)
		}
	}
	return respond, nil
}

func (repo *ProductRepository) GetById(id int) (*model.Product, error) {
	return repo.db.Search(id)
}

func (repo *ProductRepository) GetAll() ([]model.Product, error) {
	return repo.db.GetAll()
}

func (repo *ProductRepository) Save(product model.Product) error {
	return repo.db.Add(product)
}

func (repo *ProductRepository) Delete(id int) error {
	return repo.db.Delete(id)
}

func (repo *ProductRepository) Put(product model.Product) error {
	return repo.db.Update(product)
}
