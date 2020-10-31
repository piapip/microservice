package handlers

import (
	"log"
	"net/http"

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

func (p *Products) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		p.getProducts(res, req)
		return
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *Products) getProducts(res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle GET products")
	// what's that p for... it just seems very counter-intuitive. And how does it know that this is a GET...
	// So I'm suspecting that the function Encode() in ToJSON() will imprint the listProduct to the stream, in which the `res` of http.ResponseWriter resides.
	listProduct := data.GetProducts()
	err := listProduct.ToJSON(res)
	if err != nil {
		http.Error(res, "Unable to marshal json", http.StatusInternalServerError)
	}
}
