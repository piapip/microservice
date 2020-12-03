package main

import (
	"fmt"
	"testing"

	"github.com/piapip/microservice/sdk/client"
	"github.com/piapip/microservice/sdk/client/products"
)

func TestSDK(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := products.NewListProductsParams()
	products, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}

	// fmt.Printf("%#v", products.Payload[0])
	// Why is the function up there allowed too, when Get Function exists..., that's some bad design there. Anyway, either way is fine.
	fmt.Printf("%#v", products.GetPayload()[0])

	// Intentionally fail this test to print out products
	t.Fail()
}
