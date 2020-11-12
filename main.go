package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/piapip/microservice/handlers"
	"golang.org/x/net/context"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create a new serve mux and register the handlers
	serveMux := mux.NewRouter()

	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	deleteRouter := serveMux.Methods(http.MethodDelete).Subrouter()

	// create the handlers for Hello
	helloHandler := handlers.NewHello(logger)
	serveMux.Handle("/hello", helloHandler)

	// create the handlers for Goodbye
	goodbyeHandler := handlers.NewGoodbye(logger)
	serveMux.Handle("/goodbye", goodbyeHandler)

	// create the handlers for Products
	productsHandler := handlers.NewProducts(logger)
	getRouter.HandleFunc("/products", productsHandler.GetProducts)

	postRouter.HandleFunc("/products", productsHandler.AddProduct)
	postRouter.Use(productsHandler.MiddlewareProductValidation)

	putRouter.HandleFunc("/products/{id:[0-9]+}", productsHandler.UpdateProduct)
	putRouter.Use(productsHandler.MiddlewareProductValidation)

	deleteRouter.HandleFunc("/products/{id:[0-9]+}", productsHandler.DeleteProduct)
	// serveMux.Handle("/products", productsHandler)

	// SERVER CONFIGURATION
	// Customized server:
	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		logger.Println("Stating server on port ", server.Addr)
		// So normally, this http.ListenAndServe will look for handler from the default ServeMux.
		// But often enough, we want to specify exact which handler to use (for better visibility and we want to sort things our way).
		// So we assign our own server instead of using the default one.
		// The default one:
		// http.ListenAndServe(":9090", serveMux)
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
			os.Exit(1)
		}
	}()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	//won't proceed until signalChannel receive a signal
	receivedSignal := <-signalChannel

	// Gracefully Shutdown. Will stop receive request, finished all ongoing tasks then shutdown. Won't cause abruptedly disconnection.
	// So 30 seconds after shutdown command, forcefully shutdown everything, this 30-second number should be tuned.
	// shutdownContext, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))

	switch {
	case receivedSignal == os.Interrupt || receivedSignal == os.Kill:
		logger.Println("\nReceived terminate, gracefully shutdown", receivedSignal)
		shutdownContext, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancelFunc()
		server.Shutdown(shutdownContext)

	default:
		logger.Fatal("FOREIGN SIGNAL!")
	}
}
