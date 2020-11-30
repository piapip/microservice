package handlers

import (
	"net/http"

	"github.com/piapip/microservice/data"
)

// Update handles PUT requests to update products
func (p *Products) Update(res http.ResponseWriter, req *http.Request) {
	targetedProduct := req.Context().Value(KeyProduct{}).(data.Product)
	p.logger.Println("Handle PUT product id: ", targetedProduct.ID)

	err := data.UpdateProduct(targetedProduct)
	if err == data.ErrorProductNotFound {
		p.logger.Println("[ERROR] updating product not found", err)

		// http.Error(res, "Product not found", http.StatusNotFound)
		res.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in database"}, res)
		return
	}

	if err != nil {
		p.logger.Println("[ERROR] deleting record", err)

		res.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, res)
	}

	// write the no content success header
	res.WriteHeader(http.StatusNoContent)
}
