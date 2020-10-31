package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	g.l.Println("Goodbye")
	fmt.Fprintf(res, "Goodbye\n")
	res.Write([]byte("Byeeeee"))
}
