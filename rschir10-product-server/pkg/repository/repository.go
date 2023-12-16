package repository

import (


	"github.com/google/uuid"
	desc "ubulllka.com/rschir10-product-server/pkg/api"
)

var array []*desc.Product
func init() {
	array = append(array, &desc.Product{Id: "a68b823c-7ca6-44bc-b721-fb4d5312cafc", Name: "Milk", Price: 80})
	array = append(array, &desc.Product{Id: "34415c7c-f82d-4e44-88ca-ae2a1aaa92b7", Name: "Bread", Price: 40})
}



func GetAll() (*desc.ProductList, error) {
	return &desc.ProductList{
		Products: array,
	}, nil
}

func Get(in *desc.ProductId) (*desc.Product, error) {
	var product *desc.Product
	for _, e := range array {
		if e.Id == in.Id {
			product = e
			break
		}
	}
	return product, nil
}

func Insert(newProduct *desc.Product) (*desc.Product, error) {
	uuid := uuid.NewString()
	newProduct.Id = uuid
	array = append(array, newProduct)
	return newProduct, nil
}

func Update(updateProduct *desc.Product) (*desc.Product, error) {
	var product *desc.Product
	for _, e := range array {
		if e.Id == updateProduct.Id {
			e.Name = updateProduct.Name
			e.Price = updateProduct.Price
			product = e
			break
		}
	}
	return product, nil
}

func Remove(in *desc.ProductId) (*desc.Product, error) {
	var product *desc.Product
	var index int
	for i, e := range array {
		if e.Id == in.Id {
			product = e
			index = i
			break
		}
	}
	array = append(array[:index], array[index+1:]...)
	return product, nil
}