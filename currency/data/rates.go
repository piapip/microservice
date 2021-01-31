package data

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

	hclog "github.com/hashicorp/go-hclog"
)

// ExchangeRates -
type ExchangeRates struct {
	logger hclog.Logger
	rates  map[string]float64
}

// NewExchangeRates -
func NewExchangeRates(l hclog.Logger) (*ExchangeRates, error) {
	// exchangeRate := &ExchangeRates{logger: l, rate: make(map[string]float64)}
	exchangeRate := &ExchangeRates{logger: l, rates: map[string]float64{}}

	err := exchangeRate.getRates()

	return exchangeRate, err
}

// GetRate returns ratio of any 2 available currencies
// return -1 if either base or destination is not an available currency
func (e *ExchangeRates) GetRate(base string, destination string) (float64, error) {
	// check if base and destination are available currencies
	baseRate, ok := e.rates[base]
	if !ok {
		return -1, fmt.Errorf("Rate not found for currency %s", base)
	}

	destinationRate, ok := e.rates[destination]
	if !ok {
		return -1, fmt.Errorf("Rate not found for currency %s", destination)
	}

	return destinationRate / baseRate, nil
}

// getRates will pull data from online API and map it to e.rate
func (e *ExchangeRates) getRates() error {
	response, err := http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Expecting success code 200 got %d", response.StatusCode)
	}
	defer response.Body.Close()

	// Decoding
	// Initiate collection
	metadata := &Cubes{}
	// Just like how we have our ToJSON and FromJSON, golang has a lot of repetitive pattern
	xml.NewDecoder(response.Body).Decode(metadata)

	// loop over the collections, convert them to float then add them to the map
	for _, cube := range metadata.CubeData {
		rate, err := strconv.ParseFloat(cube.Rate, 64)
		if err != nil {
			return err
		}

		e.rates[cube.Currency] = rate
	}

	// hard code for the rate of EUR -> EUR
	e.rates["EUR"] = 1

	return nil
}

// Cube is for XML parsing
type Cube struct {
	// extracting currency and rate from the bottom layer of Cubes
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}

// Cubes is for XML parsing, normally in XML we will have 3 layers all with the same object (Cube in this example) with slightly different parameters, at least that's what I heard.
type Cubes struct {
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}
