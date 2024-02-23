package dto

import "github.com/Sotnasjeff/hexagonal-architecture-api/app"

type Product struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func (p *Product) Bind(product app.Product) (*app.Product, error) {
	if p.ID != "" {
		product.Id = p.ID
	}
	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status

	_, err := product.IsValid()
	if err != nil {
		return &app.Product{}, err
	}
	return &product, nil
}
