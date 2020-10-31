package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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
	case http.MethodPost:
		p.addProducts(res, req)
		return
	case http.MethodPut:
		path := req.URL.Path
		// PUT always has ID in the URL and normally, regexp is the quickest way to do but it's not easy.
		// regex takes every "/(somenumbers)" like "/2" or "/14" so let's take into consideration to not repeat this pattern later on, so this regex will still work.
		// p.l.Println(path)
		regex := regexp.MustCompile("/([0-9]+)")
		// so like I mention just above, in case we get more than 1 piece of string with that pattern, we will need to handle such situation
		// and even though I name the variable 'groupItemID', it is still an array of string with only 1 item (hopefully)
		groupItemID := regex.FindAllStringSubmatch(path, -1)
		if len(groupItemID) != 1 {
			p.l.Println("Invalid URL not 1 id")
			http.Error(res, "Invalid URL", http.StatusBadRequest)
			return
		}
		// ... overall, this kind of defensive coding methods is painful, don't do this
		if len(groupItemID[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(res, "Invalid URL", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(groupItemID[0][1])
		if err != nil {
			p.l.Println("Invalid URI unable to convert to number", groupItemID[0][1])
			http.Error(res, "Invalid URL, bad ID", http.StatusBadRequest)
		}
		p.updateProduct(id, res, req)
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

func (p *Products) addProducts(res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle POST products")

	newProduct := &data.Product{}
	err := newProduct.FromJSON(req.Body)
	if err != nil {
		http.Error(res, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(newProduct)
}

func (p *Products) updateProduct(id int, res http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle PUT products")

	newProduct := &data.Product{}

	err := newProduct.FromJSON(req.Body)
	if err != nil {
		http.Error(res, "Something's wrong with the server. Unable to convert from JSON to Product struct", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, newProduct)
	if err == data.ErrorProductNotFound {
		http.Error(res, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(res, "Product not found", http.StatusInternalServerError)
		return
	}
}
