package main

import (
	"net"
	"os"

	go_hclog "github.com/hashicorp/go-hclog"
	protoS "github.com/piapip/microservice/currency/protoS/currency"
	"github.com/piapip/microservice/currency/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger := go_hclog.Default()

	// Prepare to create server
	currencyServer := server.NewCurrency(logger)

	// create a new gRPC server, use WithInsecure to allow http connections
	grpcServer := grpc.NewServer()

	// Create
	protoS.RegisterCurrencyServer(grpcServer, currencyServer)

	// need this to test with grpcurl, we might want to disable this line in production though
	reflection.Register(grpcServer)

	// Start the server
	// grpcServer has a method called "Serve". Serve() is kinda similar to ListenAndServe in normal Go.
	// The difference is we have to specify a "net listener" for grpc server.

	// Get a net listener, create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		logger.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	// listen for requests
	grpcServer.Serve(l)
}
