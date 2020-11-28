package data

import (
	"fmt"
	"time"
)

// ErrorProductNotFound is returned when you can't find any wanted product.
var ErrorProductNotFound = fmt.Errorf("Product not found")

// Product for some goods.
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"` //customized validation sample
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products is a list of available products
type Products []*Product

// GetProducts returns the products from the data store
func GetProducts() Products {
	return productsList
}

// GetProductByID returns a product which matches the id from the list.
// Returns ErrorProductNotFound if there's no matching result.
func GetProductByID(id int) (*Product, error) {
	index := findIndexByProductID(id)
	if id == -1 {
		return nil, ErrorProductNotFound
	}

	return productsList[index], nil
}

// AddProduct will add a product to the list.
func AddProduct(p Product) {
	p.ID = getNextID()
	productsList = append(productsList, &p)
}

// UpdateProduct will update the selected product in the list.
// Returns ErrorProductNotFound if there's no matching result.
func UpdateProduct(p Product) error {
	index := findIndexByProductID(p.ID)
	if index == -1 {
		return ErrorProductNotFound
	}

	productsList[index] = &p
	return nil
}

// DeleteProduct will delete the product with respective id in the list.
// Returns ErrorProductNotFound if there's no matching result.
func DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return ErrorProductNotFound
	}

	// productsList = append(productsList[:i], productsList[i+1]...)
	productsList = append(productsList[i:], productsList[:i+1]...)
	return nil
}

func getNextID() int {
	lastProduct := productsList[len(productsList)-1]
	return lastProduct.ID + 1
}

// // findIndex finds the index of a product in the database
// // returns -1 when no product can be found
// func findProduct(id int) (*Product, int, error) {
// 	for index, p := range productsList {
// 		if p.ID == id {
// 			return p, index, nil
// 		}
// 	}

// 	return nil, -1, ErrorProductNotFound
// }

func findIndexByProductID(id int) int {
	for index, product := range productsList {
		if product.ID == id {
			return index
		}
	}

	return -1
}

// productsList is a list of coffee stuff. Should be deleted and should never be initiated this way.
var productsList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "qwe-rty-uio",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "asd-ghj-klm",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
