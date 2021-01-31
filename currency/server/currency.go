package server

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/piapip/microservice/currency/data"
	protoS "github.com/piapip/microservice/currency/protoS/currency"
)

// Currency is a gRPC server it implements the methods defined by the CurrencyServer interface
type Currency struct {
	protoS.UnimplementedCurrencyServer
	logger hclog.Logger
	rates  *data.ExchangeRates
}

// NewCurrency creates a new Currency server
func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	return &Currency{rates: r, logger: l}
}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate
// for the two given currencies.
func (c *Currency) GetRate(ctx context.Context, rr *protoS.RateRequest) (*protoS.RateResponse, error) {
	c.logger.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err

	}

	return &protoS.RateResponse{Rate: rate}, nil
}
