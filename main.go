package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.Println("Hello world")
		data, err := ioutil.ReadAll(req.Body)

		if err != nil {
			http.Error(res, "Oops", http.StatusBadRequest)
			return
		}
		log.Printf("Data: %s\n", data)
		fmt.Fprintf(res, "Hello %s", data)
	})

	http.HandleFunc("/goodbye", func(res http.ResponseWriter, rq *http.Request) {
		log.Println("Goodbye")
	})

	http.ListenAndServe(":9090", nil)
}
