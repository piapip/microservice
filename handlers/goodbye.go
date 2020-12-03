package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// Goodbye handlers will return an error for domain which is under development
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye returns a Goodbye object as dependency injection (?)
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	g.l.Println("Goodbye")
	fmt.Fprintf(res, "Goodbye\n")
	res.Write([]byte("Byeeeee"))
}
