package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello handlers will return an error for domain which is under development
type Hello struct {
	l *log.Logger
}

// NewHello returns a Hello object as dependency injection (?)
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	h.l.Println("Hello world")
	data, err := ioutil.ReadAll(req.Body)

	if err != nil {
		http.Error(res, "Oops", http.StatusBadRequest)
		return
	}
	log.Printf("Data: %s\n", data)
	fmt.Fprintf(res, "Hello %s", data)
}
