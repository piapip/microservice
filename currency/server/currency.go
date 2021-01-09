package server

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	protoS "github.com/piapip/microservice/currency/protoS/currency"
)

// Currency - ...
type Currency struct {
	protoS.UnimplementedCurrencyServer
	logger hclog.Logger
}

// NewCurrency creates a new Currency server
func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{logger: l}
}

// GetRate - ...
func (c *Currency) GetRate(ctx context.Context, rr *protoS.RateRequest) (*protoS.RateResponse, error) {
	c.logger.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	return &protoS.RateResponse{Rate: 0.5}, nil
}
