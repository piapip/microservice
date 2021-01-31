package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	hclog "github.com/hashicorp/go-hclog"
)

// Hello handlers will return an error for domain which is under development
type Hello struct {
	l hclog.Logger
}

// NewHello returns a Hello object as dependency injection (?)
func NewHello(l hclog.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	h.l.Debug("Hello world")
	data, err := ioutil.ReadAll(req.Body)

	if err != nil {
		http.Error(res, "Oops", http.StatusBadRequest)
		return
	}
	// log.Printf("Data: %s\n", data)
	fmt.Fprintf(res, "Hello %s", data)
}
