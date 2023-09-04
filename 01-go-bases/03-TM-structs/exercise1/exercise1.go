package main

import (
	"fmt"
)

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products = []Product{{
	ID:          0,
	Name:        "Calculator 154648",
	Price:       100,
	Description: "",
	Category:    "Calculator",
}, {
	ID:          1,
	Name:        "Calculator 45465",
	Price:       200.5,
	Description: "",
	Category:    "Calculator",
}}

func (p Product) Save() {
	Products = append(Products, p)
}

func GetAll() {
	fmt.Println(Products)
}

func getById(id int) (Product, bool) {
	for _, p := range Products {
		if p.ID == id {
			return p, true
		}
	}
	return Product{}, false
}

func main() {

	newProduct := Product{
		ID:          2,
		Name:        "Calculator 304648",
		Price:       100.50,
		Description: "",
		Category:    "Calculator",
	}
	newProduct.Save()
	p, ok := getById(2)
	if ok {
		fmt.Println(p)
	}
	GetAll()
}
