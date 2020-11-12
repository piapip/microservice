package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
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

//Validate check if the product information is good enough
func (p *Product) Validate() error {
	validate := validator.New()
	// tag name and the validation function
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}

// Products is a list of available products
type Products []*Product

// ToJSON will convert list of product to JSON format
func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	// All kind of funny shenanigan before encoding will go here. Like I'll only encode those goods that have a SKU.
	// Doing something with p here.
	return encoder.Encode(p)
}

// FromJSON converts items in JSON format from the stream to Product struct
func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	// This one is quite confusing tbh.
	return decoder.Decode(p)
}

// GetProducts returns the products from the data store
func GetProducts() Products {
	return productsList
}

// AddProduct will add a product to the list.
func AddProduct(p *Product) {
	p.ID = getNextID()
	productsList = append(productsList, p)
}

// UpdateProduct will update the selected product in the list.
func UpdateProduct(id int, p *Product) error {
	_, index, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productsList[index] = p
	return nil
}

func getNextID() int {
	lastProduct := productsList[len(productsList)-1]
	return lastProduct.ID + 1
}

func findProduct(id int) (*Product, int, error) {
	for index, p := range productsList {
		if p.ID == id {
			return p, index, nil
		}
	}

	return nil, -1, ErrorProductNotFound
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
