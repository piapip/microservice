package handlers

import (
	"net/http"

	"github.com/piapip/microservice/data"
)

// swagger:route POST /products products createProduct
//
// Create a new product
// responses:
// 		200: productResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Products) Create(res http.ResponseWriter, req *http.Request) {
	// fetch the product from the context
	p.logger.Println("Handle POST products")

	// Mind you the struct, it depends on the context, sometimes it's .(data.Product), sometimes it's .(*data.Product).
	// It should be somewhere around the middleware files.
	newProduct := req.Context().Value(KeyProduct{}).(data.Product)
	p.logger.Printf("[DEBUG] inserting new product: %#v\n", newProduct)
	data.AddProduct(newProduct)
}
