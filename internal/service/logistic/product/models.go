package product

import "fmt"

type Product struct {
	ID    int     `json:"id"`
	Title string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *Product) String() string {
	return fmt.Sprintf("Product:\nID:%d\nTitle:%s\nPrice:%.2f", p.ID, p.Title, p.Price)
}
