package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float32 `json:"price"`
	SKU           string  `json:"sku"`
	CreatedOn     string  `json:"-"`
	UpdatedOn     string  `json:"-"`
	DeletedOn     string  `json:"-"`
	PercentageOff int     `json:"percentageOff"`
}

//list of all products
type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

//single product GET
func GetSingleProduct(ID int) *Product {
	fmt.Println("getting single product")
	fetchedProduct := productList[ID]
	return fetchedProduct
}

func (p *Product) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

var productList = []*Product{
	&Product{
		ID:            1,
		Name:          "Latte",
		Description:   "Frothy milky coffee.",
		Price:         2.45,
		SKU:           "10001",
		CreatedOn:     time.Now().UTC().String(),
		UpdatedOn:     time.Now().UTC().String(),
		PercentageOff: 100,
	},
	&Product{
		ID:            2,
		Name:          "Espresso",
		Description:   "Short, strong coffee with no milk.",
		Price:         3.45,
		SKU:           "10002",
		CreatedOn:     time.Now().UTC().String(),
		UpdatedOn:     time.Now().UTC().String(),
		PercentageOff: 100,
	},
	&Product{
		ID:            3,
		Name:          "Flat White",
		Description:   "Very milky.",
		Price:         2.45,
		SKU:           "10003",
		CreatedOn:     time.Now().UTC().String(),
		UpdatedOn:     time.Now().UTC().String(),
		PercentageOff: 100,
	},
	&Product{
		ID:            4,
		Name:          "Mocha",
		Description:   "Coffee with chocolate. What's better.",
		Price:         3.15,
		SKU:           "10004",
		CreatedOn:     time.Now().UTC().String(),
		UpdatedOn:     time.Now().UTC().String(),
		PercentageOff: 100,
	},
}
