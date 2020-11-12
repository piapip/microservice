package data

import "testing"

func TestCheckValidations(t *testing.T) {
	p := &Product{
		Name:  "Test",
		Price: 20,
		SKU:   "abc-def-ikl",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
