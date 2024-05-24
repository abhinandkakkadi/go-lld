package main

import (
	"time"

	"github.com/abhinandkakkadi/go-lld/factory-method/products"
)

func main() {
	// New method is usually used for creating objects in go
	product := products.New("phone")
	_ = product

	// Commonly used factory method in Go
	product2 := products.Product{
		Name:      "laptop",
		CreatedAt: time.Now(),
	}
	_ = product2
}
