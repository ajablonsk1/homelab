package parse

import (
	"fmt"
	"strconv"
)

type Product struct {
	Name  string `redis:"name"`
	Month string `redis:"month"`
	URL   string `redis:"url"`
	Cost  int    `redis:"cost"`
}

func (p *Product) Update(product *Product) {
	if product.Name != "" {
		p.Name = product.Name
	}
	if product.Cost != 0 {
		p.Cost = product.Cost
	}
	if product.Month != "" {
		p.Month = product.Month
	}
	if product.URL != "" {
		p.URL = product.URL
	}
}

func ProductFromMap(m map[string]string) (Product, error) {
	cost, err := strconv.Atoi(m["cost"])
	if err != nil {
		return Product{}, fmt.Errorf("encountered error during converting cost %s to int: %s", m["cost"], err)
	}

	return Product{
		Name:  m["name"],
		Month: m["month"],
		URL:   m["url"],
		Cost:  cost,
	}, nil
}
