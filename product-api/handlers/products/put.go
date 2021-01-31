package handlers

import (
	"net/http"

	"github.com/piapip/microservice/product-api/data"
)

// swagger:route PUT /products products updateProduct
//
// Update a product details
// responses:
//  	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update products
func (p *Products) Update(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")

	// fetch the product from the context
	targetedProduct := req.Context().Value(KeyProduct{}).(data.Product)
	p.logger.Debug("Handle PUT product", "id", targetedProduct.ID)

	err := p.productDB.UpdateProduct(targetedProduct)
	if err == data.ErrorProductNotFound {
		p.logger.Error("Unable to update product not found", "error", err)

		// http.Error(res, "Product not found", http.StatusNotFound)
		res.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in database"}, res)
		return
	}

	if err != nil {
		p.logger.Error("Unable to delete record", "error", err)

		res.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, res)
	}

	// write the no content success header
	res.WriteHeader(http.StatusNoContent)
}
