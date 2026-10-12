package parse

import (
	"flag"
	"fmt"
	"slices"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var months = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

func AddProductArguments() (Product, error) {
	name, cost, month, url := getProductFlags()

	if name == "" {
		return Product{}, fmt.Errorf("product name is required")
	}

	caser := cases.Title(language.English)
	month = caser.String(month)
	if month != "" && !slices.Contains(months, month) {
		return Product{}, fmt.Errorf("invalid month %s", month)
	}

	return Product{Name: name, Cost: cost, Month: month, URL: url}, nil
}

func UpdateProductArguments() (Product, error) {
	product, err := AddProductArguments()
	return product, err
}

func RemoveProductArguments() (string, error) {
	name, _, _, _ := getProductFlags()
	return name, nil
}

func getProductFlags() (string, int, string, string) {
	name := *flag.String("name", "", "Product name")
	cost := *flag.Int("cost", 0, "Product cost")
	month := *flag.String("month", "", "Month to buy")
	url := *flag.String("url", "", "Product URL")

	return name, cost, month, url
}
