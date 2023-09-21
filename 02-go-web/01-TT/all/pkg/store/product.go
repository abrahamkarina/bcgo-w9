package store

import (
	"encoding/json"
	"fmt"
	"github.com/abrahamkarina/bcgo-w9/02-go-web/01-TT/exercise1n2/internal/domain/model"
	"os"
)

type Store interface {
	Init(filename string) error
	Search(id int) (*model.Product, error)
	GetAll() ([]model.Product, error)
	Delete(id int) error
	Add(product model.Product) error
	Update(product model.Product) error
}

type store struct {
	filename string
	products []model.Product
}

func NewStore() Store {
	return &store{}
}

func (s *store) Init(filename string) error {
	s.filename = filename

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		s.products = []model.Product{}
		return s.saveToFile()
	} else if err != nil {
		return err
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &s.products)
}

func (s *store) Search(id int) (*model.Product, error) {
	for _, product := range s.products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, nil
}

func (s *store) Delete(id int) error {
	index, err := s.searchIndexById(id)
	if err != nil {
		return err
	}

	s.products = append(s.products[:index], s.products[index+1:]...)
	return s.saveToFile()
}

func (s *store) searchIndexById(id int) (int, error) {
	index := -1
	for i, product := range s.products {
		if product.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return 0, fmt.Errorf("product not found with ID: %d", id)
	}
	return index, nil
}

func (s *store) GetAll() ([]model.Product, error) {
	if s.products == nil {
		return nil, fmt.Errorf("products not found")
	}
	products := make([]model.Product, 0)
	for _, p := range s.products {
		products = append(products, p)
	}
	return products, nil
}

func (s *store) Update(product model.Product) error {
	index, err := s.searchIndexById(product.ID)
	if err != nil {
		return err
	}
	s.products[index] = product
	return s.saveToFile()
}

func (s *store) Add(product model.Product) error {
	s.products = append(s.products, product)
	return s.saveToFile()
}

func (s *store) saveToFile() error {
	data, err := json.MarshalIndent(s.products, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.filename, data, 0644)
}
