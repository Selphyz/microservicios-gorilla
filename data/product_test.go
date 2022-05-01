package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Cosas",
		Price: 1.00,
		SKU:   "abs-afd-aqwe",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
