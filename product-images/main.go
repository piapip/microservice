package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	hclog "github.com/hashicorp/go-hclog"
)

func main() {

	// logger := log.New(os.Stdout, "product-images", log.LstdFlags)
	// why this logger?
	logger := hclog.New(
		&hclog.LoggerOptions{
			Name:  "product-images",
			Level: hclog.LevelFromString("debug"),
		},
	)

	// create a logger for the server from the default logger
	serverLogger := logger.StandardLogger(&hclog.StandardLoggerOptions{InferLevels: true})

	serveMux := mux.NewRouter()

	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
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
