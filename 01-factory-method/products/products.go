package products

import "time"

type Product struct {
	Name      string
	CreatedAt time.Time
}

func New(name string) *Product {
	return &Product{
		Name:      name,
		CreatedAt: time.Now(),
	}
}
