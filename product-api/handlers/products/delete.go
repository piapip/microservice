package handlers

import (
	"net/http"

	"github.com/piapip/microservice/product-api/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
//
// Delete a product with the respective ID.
// responses:
// 		201: noContentResponse
// 	404: errorResponse
//  500: errorResponse

// Delete handles DELETE requests and removes items from the database
func (p *Products) Delete(res http.ResponseWriter, rq *http.Request) {
	id := getProductID(rq)

	p.logger.Debug("Delete product", "id", id)

	err := p.productDB.DeleteProduct(id)
	if err == data.ErrorProductNotFound {
		p.logger.Error("Unable to delete record", "error", err)

		// http.Error(res, "Product not found", http.StatusNotFound)
		res.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, res)
		return
	}

	if err != nil {
		p.logger.Error("Unable to delete record", "error", err)

		res.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, res)
	}
}
