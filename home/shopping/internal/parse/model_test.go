package parse

import "testing"

func TestUpdateProduct(t *testing.T) {
	product := Product{Name: "Example Product", Month: "January", URL: "example.com", Cost: 200}
	productCopy := product

	products := []Product{
		{},
		{Name: "New Name"},
		{Month: "February"},
		{URL: "new.com"},
		{Cost: 150},
		{Name: "New Name", Month: "July"},
		{Name: "New Name", Cost: 30},
		{Name: "New Name", Month: "July", URL: "new.com"},
		{Name: "New Name", Month: "July", URL: "new.com", Cost: 30},
	}

	expectedResults := []Product{
		{Name: "Example Product", Month: "January", URL: "example.com", Cost: 200},
		{Name: "New Name", Month: "January", URL: "example.com", Cost: 200},
		{Name: "Example Product", Month: "February", URL: "example.com", Cost: 200},
		{Name: "Example Product", Month: "January", URL: "new.com", Cost: 200},
		{Name: "Example Product", Month: "January", URL: "example.com", Cost: 150},
		{Name: "New Name", Month: "July", URL: "example.com", Cost: 200},
		{Name: "New Name", Month: "January", URL: "example.com", Cost: 30},
		{Name: "New Name", Month: "July", URL: "new.com", Cost: 200},
		{Name: "New Name", Month: "July", URL: "new.com", Cost: 30},
	}

	for i, p := range products {
		product.Update(&p)
		if product != expectedResults[i] {
			t.Errorf("updated product %+v doesn't match expected one %+v\n", product, expectedResults[i])
		}
		product = productCopy
	}
}
