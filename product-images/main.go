package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	gorilla_handlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	hclog "github.com/hashicorp/go-hclog"
	"github.com/piapip/microservice/product-images/files"
	"github.com/piapip/microservice/product-images/handlers"
)

var basePath = "C:/Users/thovi/Desktop/test/microservice/product-images/imagestore"

func main() {

	// logger := log.New(os.Stdout, "product-images", log.LstdFlags)
	// why this logger?
	logger := hclog.New(
		&hclog.LoggerOptions{
			Name:  "product-images",
			Level: hclog.LevelFromString("debug"),
		},
	)

	// create the storage class, using local storage
	// max file size: 5MB
	storage, err := files.NewLocal(basePath, 1024*1000*5)
	if err != nil {
		logger.Error("Unable to create storage", "error", err)
		os.Exit(1)
	}

	// create new serveMux and register the handlers
	serveMux := mux.NewRouter()

	// create files handlers
	fileHandlers := handlers.NewFiles(storage, logger)

	// filename regex: {filename:[a-zA-Z]+\\.[a-z]{3}} <- this regex is pretty stupid but it's good enough for now.
	// A little bit difference compared to product-api
	postHandler := serveMux.Methods(http.MethodPost).Subrouter()
	postHandler.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fileHandlers.UploadREST)
	postHandler.HandleFunc("/", fileHandlers.UploadMultipart)

	// When I call the GET, please give me the picture, it will find the path given from http.Dir,
	// then it tracks down the path based on the URL, this case is: basePath/images/1/test.jpg
	// however, in imagestore, we don't have any directory called images, so we use the http.StripPrefix like below to tell the handler to ignore /images/
	// It's specially built for this purpose so don't ask too much about it.
	getHandler := serveMux.Methods(http.MethodGet).Subrouter()
	getHandler.Handle("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", http.StripPrefix("/images/", http.FileServer(http.Dir(basePath))))

	// create a logger for the server from the default logger
	serverLogger := logger.StandardLogger(&hclog.StandardLoggerOptions{InferLevels: true})
	corsHandler := gorilla_handlers.CORS(gorilla_handlers.AllowedOrigins([]string{"http://localhost:3000"}))

	server := http.Server{
		Addr:         ":9091",
		Handler:      corsHandler(serveMux),
		ErrorLog:     serverLogger,
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive

	}

	go func() {
		// logger.Println("Stating server on port ", server.Addr)
		logger.Info("Starting server", "bind_address", server.Addr)

		err := server.ListenAndServe()
		if err != nil {
			// logger.Fatal(err)
			logger.Error("Unable to start server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	//won't proceed until signalChannel receive a signal
	receivedSignal := <-signalChannel

	// Gracefully Shutdown.
	switch {
	case receivedSignal == os.Interrupt || receivedSignal == os.Kill:
		logger.Info("\nReceived terminate, gracefully shutdown", "signal", receivedSignal)
		shutdownContext, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancelFunc()
		server.Shutdown(shutdownContext)

	default:
		logger.Error("FOREIGN SIGNAL!")
	}
}
