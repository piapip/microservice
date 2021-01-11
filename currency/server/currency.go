package server

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	protoS "github.com/piapip/microservice/currency/protoS/currency"
)

// Currency is a gRPC server it implements the methods defined by the CurrencyServer interface
type Currency struct {
	protoS.UnimplementedCurrencyServer
	logger hclog.Logger
}

// NewCurrency creates a new Currency server
func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{logger: l}
}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate
// for the two given currencies.
func (c *Currency) GetRate(ctx context.Context, rr *protoS.RateRequest) (*protoS.RateResponse, error) {
	c.logger.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	return &protoS.RateResponse{Rate: 0.5}, nil
}
