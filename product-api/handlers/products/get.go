package handlers

import (
	"net/http"

	"github.com/piapip/microservice/product-api/data"
)

// swagger:route GET /products products listProducts
//
// Return a list of products from the database
// responses:
//  200: productsResponse

// ListAll handles GET requests and returns all item in the list
func (p *Products) ListAll(res http.ResponseWriter, rq *http.Request) {
	p.logger.Debug("Get all records")

	// We are returning text/plain here.
	// Here's why: if we don't specify the content type, some functions in the auto gen codes (down side of using auto gen code) returns text/plain when it's provided a string.
	// I'm not sure which function is that I need to do this debug again. And yes, I'm lazy sue me.
	// That function is Submit(), ct (parsed media type) is expecting text/plain when it's supposed to expect application/json instead.
	// What is media type?
	// CONFLICT!!! Have to solve it here.
	res.Header().Add("Content-Type", "application/json")

	products, err := p.productDB.GetProducts("")
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, res)
		return
	}

	err = data.ToJSON(products, res)
	if err != nil {
		// we should never be here but log the error just incase
		p.logger.Error("Unable to serialize product", "error", err)
	}
}

// swagger:route GET /products/{id} products listSingleProduct
//
// Return a product with the respective ID in the database
// responses:
//  200: productResponse
//  404: errorResponse

// ListSingle handles GET requests, return the product with chosen ID.
func (p *Products) ListSingle(res http.ResponseWriter, rq *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	id := getProductID(rq)

	p.logger.Debug("Get record", "id", id)

	product, err := p.productDB.GetProductByID(id, "")

	switch err {
	case nil:
	case data.ErrorProductNotFound:
		p.logger.Error("Unable to fetch product", "error", err)
		res.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, res)
		return
	default:
		p.logger.Error("Unable to fetch product", "error", err)
		res.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, res)
		return
	}

	err = data.ToJSON(product, res)
	if err != nil {
		// should never happen, unless someone try to be funny
		http.Error(res, "Unable to marshal json", http.StatusInternalServerError)
	}
}
