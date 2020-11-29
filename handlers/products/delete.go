package handlers

import (
	"net/http"

	"github.com/piapip/microservice/data"
)

// Delete handles DELETE requests and removes items from the database
func (p *Products) Delete(res http.ResponseWriter, rq *http.Request) {
	id := getProductID(rq)

	p.logger.Println("[DEBUG] delete product id: ", id)

	err := data.DeleteProduct(id)
	if err == data.ErrorProductNotFound {
		p.logger.Println("[ERROR] deleting record id does not exist")

		// http.Error(res, "Product not found", http.StatusNotFound)
		res.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, res)
		return
	}

	if err != nil {
		p.logger.Println("[ERROR] deleting record", err)

		res.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, res)
	}
}
