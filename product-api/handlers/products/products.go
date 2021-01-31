package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	hclog "github.com/hashicorp/go-hclog"
	"github.com/piapip/microservice/product-api/data"
)

// KeyProduct is a key used for the Product object in the context
type KeyProduct struct{}

// Products handler for getting and updating products
type Products struct {
	logger    hclog.Logger
	validator *data.Validation
	productDB *data.ProductsDB
}

// NewProducts returns a new products handler with the given logger
func NewProducts(logger hclog.Logger, validator *data.Validation, productDB *data.ProductsDB) *Products {
	return &Products{logger, validator, productDB}
}

// ErrorInvalidProductPath is an error message when the product path is not valid
var ErrorInvalidProductPath = fmt.Errorf("Invalid path, path should be :/products/[id]")

// GenericError is a generic error message returned by the server
// don't write with a space like this [json: "message"] it will shoot error
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getProductID(req *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(req)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen, unless someone try to be funny
		panic(err)
	}

	return id
}
