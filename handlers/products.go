package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/piapip/microservice/data"
)

// Products is a http.handler
type Products struct {
	l *log.Logger
}

// NewProducts returns a Products object as dependency injection (?)
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// GetProducts returns the products from the data store
func (p *Products) GetProducts(res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle GET products")
	// what's that p for... it just seems very counter-intuitive. And how does it know that this is a GET...
	// So I'm suspecting that the function Encode() in ToJSON() will imprint the listProduct to the stream, in which the `res` of http.ResponseWriter resides.
	listProduct := data.GetProducts()

	err := listProduct.ToJSON(res)
	if err != nil {
		http.Error(res, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// AddProduct adds a products
func (p *Products) AddProduct(res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle POST products")

	newProduct := req.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&newProduct)
}

// UpdateProduct does updating stuffs
func (p *Products) UpdateProduct(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(res, "Unable to convert id", http.StatusBadRequest)
	}

	p.l.Println("Handle PUT product ", id)
	newProduct := req.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &newProduct)
	if err == data.ErrorProductNotFound {
		http.Error(res, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(res, "Product not found", http.StatusInternalServerError)
		return
	}
}

// KeyProduct is a key type for Product in context
type KeyProduct struct{}

// MiddlewareProductValidation will extract the product struct, fill the empty struct with data then put it into req's Context() with Value of sampleProd and key of KeyProduct{}
func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// if you notice, we assign sampleProd := &models.Product{} with & because we will use it straight up in the function above
		// models.AddProduct(&sampleProd) used to be models.AddProduct(sampleProd) without the "&"
		newProduct := data.Product{}

		// extracting data from req to JSON
		err := newProduct.FromJSON(req.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing products", err)
			http.Error(res, "Something's wrong with the server. Unable to convert from JSON to Product struct", http.StatusBadRequest)
			return
		}

		// validate the product
		err = newProduct.Validate()
		if err != nil {
			p.l.Println("[ERROR] false product", err)
			http.Error(res, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		// add the product to the context
		ctx := context.WithValue(req.Context(), KeyProduct{}, newProduct)
		req = req.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(res, req)
	})
}
