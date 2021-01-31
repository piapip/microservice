package handlers

import (
	"fmt"
	"net/http"

	hclog "github.com/hashicorp/go-hclog"
)

// Goodbye handlers will return an error for domain which is under development
type Goodbye struct {
	l hclog.Logger
}

// NewGoodbye returns a Goodbye object as dependency injection (?)
func NewGoodbye(l hclog.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	g.l.Debug("Goodbye")
	fmt.Fprintf(res, "Goodbye\n")
	res.Write([]byte("Byeeeee"))
}
