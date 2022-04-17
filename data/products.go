package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID           int			`json:"id"`
	Name         string		`json:"name"`
	Description  string		`json:"description"`
	Price        float32	`json:"price"`
	SKU          string		`json:"sku"`
	CreatedOn    string		`json:"-"`
	DeletedOn    string		`json:"-"`
	UpdatedOn    string		`json:"-"`
}
type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
func GetProducts() Products {
	return productList
}
var productList = []*Product{
	{
		ID:            1,
		Name:          "Latte",
		Description:   "Cafe con leche",
		Price:         2.45,
		SKU:					 "asd123",
		CreatedOn: 		 time.Now().UTC().String(),
		UpdatedOn: 		 time.Now().UTC().String(),	
	},
	{
		ID:            2,
		Name:          "Espresso",
		Description:   "Cafe corto y fuerte",
		Price:         1.99,
		SKU:					 "zxczx2",
		CreatedOn: 		 time.Now().UTC().String(),
		UpdatedOn: 		 time.Now().UTC().String(),	
	},
}