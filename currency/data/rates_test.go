package data

import (
	"fmt"
	"testing"

	hclog "github.com/hashicorp/go-hclog"
)

func TestNewExchangeRates(t *testing.T) {
	tr, err := NewExchangeRates(hclog.Default())

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Rates: %#v", tr.rates)
}
