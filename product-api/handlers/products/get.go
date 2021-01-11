package handlers

import (
	"context"
	"net/http"

	protoServer "github.com/piapip/microservice/currency/protoS/currency"
	"github.com/piapip/microservice/data"
)

// swagger:route GET /products products listProducts
//
// Return a list of products from the database
// responses:
//  200: productsResponse

// ListAll handles GET requests and returns all item in the list
func (p *Products) ListAll(res http.ResponseWriter, rq *http.Request) {
	p.logger.Println("[DEBUG] get all records")

	// We are returning text/plain here.
	// Here's why: if we don't specify the content type, some functions in the auto gen codes (down side of using auto gen code) returns text/plain when it's provided a string.
	// I'm not sure which function is that I need to do this debug again. And yes, I'm lazy sue me.
	// That function is Submit(), ct (parsed media type) is expecting text/plain when it's supposed to expect application/json instead.
	// What is media type?
	// CONFLICT!!! Have to solve it here.
	res.Header().Add("Content-Type", "application/json")

	products := data.GetProducts()

	err := data.ToJSON(products, res)
	if err != nil {
		// we should never be here but log the error just incase
		p.logger.Println("[ERROR] serializing product", err)
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
	id := getProductID(rq)

	p.logger.Println("[DEBUG] get record id", id)

	product, err := data.GetProductByID(id)

	switch err {
	case nil:
	case data.ErrorProductNotFound:
		p.logger.Println("[ERROR] fetching product", err)
		res.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, res)
		return
	default:
		p.logger.Println("[ERROR] fetching product", err)
		res.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, res)
		return
	}

	// get exchange rate from currency client
	rateRequest := &protoServer.RateRequest{
		Base:        protoServer.Currencies(protoServer.Currencies_value["EUR"]),
		Destination: protoServer.Currencies(protoServer.Currencies_value["GBP"]),
	}

	response, err := p.currencyClient.GetRate(context.Background(), rateRequest)
	if err != nil {
		p.logger.Println("[Error] error getting new rate", err)
		data.ToJSON(&GenericError{Message: err.Error()}, res)
	}

	// p.logger.Printf("Resp: %#v", response)

	product.Price = product.Price * response.Rate

	err = data.ToJSON(product, res)
	if err != nil {
		// should never happen, unless someone try to be funny
		http.Error(res, "Unable to marshal json", http.StatusInternalServerError)
	}
}
