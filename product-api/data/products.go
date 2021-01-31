package data

import (
	"context"
	"fmt"

	hclog "github.com/hashicorp/go-hclog"
	protoServer "github.com/piapip/microservice/currency/protoS/currency"
)

// ErrorProductNotFound is returned when you can't find any wanted product.
var ErrorProductNotFound = fmt.Errorf("Product not found")

// Product for some goods.
// swagger:model
type Product struct {
	// the id for this product
	//
	// required: true
	// min: 1
	ID int `json:"id"`

	//  the name for this product
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this product
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for this product
	//
	// required: true
	// min: 0.01
	// FIX TO float64
	Price float64 `json:"price" validate:"required,gt=0"`

	// the SKU for this product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	// example: abc-def-ghk
	SKU string `json:"sku" validate:"required,sku"` //customized validation sample
}

// Products is a list of available products
type Products []*Product

// ProductsDB - ...
type ProductsDB struct {
	logger   hclog.Logger
	currency protoServer.CurrencyClient
}

// NewProductsDB create a new instance of ProductsDB
func NewProductsDB(c protoServer.CurrencyClient, l hclog.Logger) *ProductsDB {
	return &ProductsDB{logger: l, currency: c}
}

// GetProducts returns the products from the data store
func (p *ProductsDB) GetProducts(currency string) (Products, error) {
	// "" means EUR
	if currency == "" {
		return productsList, nil
	}

	rate, err := p.getRate(currency)
	if err != nil {
		p.logger.Error("Unable to get rate", "currency", currency, "error", err)
		// data.ToJSON(&GenericError{Message: err.Error()}, res)
		return nil, err
	}

	productReturn := Products{}
	for _, product := range productsList {
		newProduct := *product
		newProduct.Price = newProduct.Price * rate
		productReturn = append(productReturn, &newProduct)
	}

	return productReturn, nil
}

// GetProductByID returns a product which matches the id from the list.
// Returns ErrorProductNotFound if there's no matching result.
func (p *ProductsDB) GetProductByID(id int, currency string) (*Product, error) {
	index := findIndexByProductID(id)
	if id == -1 {
		return nil, ErrorProductNotFound
	}

	if currency == "" {
		return productsList[index], nil
	}

	rate, err := p.getRate(currency)
	if err != nil {
		p.logger.Error("Unable to get rate", "currency", currency, "error", err)
		// data.ToJSON(&GenericError{Message: err.Error()}, res)
		return nil, err
	}

	newProduct := *productsList[index]
	newProduct.Price = newProduct.Price * rate
	return &newProduct, nil
}

// AddProduct will add a product to the list.
func (p *ProductsDB) AddProduct(pr Product) {
	pr.ID = getNextID()
	productsList = append(productsList, &pr)
}

// UpdateProduct will update the selected product in the list.
// Returns ErrorProductNotFound if there's no matching result.
func (p *ProductsDB) UpdateProduct(pr Product) error {
	index := findIndexByProductID(pr.ID)
	if index == -1 {
		return ErrorProductNotFound
	}

	productsList[index] = &pr
	return nil
}

// DeleteProduct will delete the product with respective id in the list.
// Returns ErrorProductNotFound if there's no matching result.
func (p *ProductsDB) DeleteProduct(id int) error {
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

func findIndexByProductID(id int) int {
	for index, product := range productsList {
		if product.ID == id {
			return index
		}
	}

	return -1
}

func (p *ProductsDB) getRate(destination string) (float64, error) {
	rateRequest := &protoServer.RateRequest{
		Base:        protoServer.Currencies(protoServer.Currencies_value["EUR"]),
		Destination: protoServer.Currencies(protoServer.Currencies_value[destination]),
	}

	response, err := p.currency.GetRate(context.Background(), rateRequest)

	// p.logger.Printf("Resp: %#v", response)

	return response.Rate, err
}

// productsList is a list of coffee stuff. Should be deleted and should never be initiated this way.
var productsList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "qwe-rty-uio",
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "asd-ghj-klm",
	},
}
