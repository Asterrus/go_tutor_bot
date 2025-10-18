package product

type Product struct {
	ID    int     `json:"id"`
	Title string  `json:"name"`
	Price float64 `json:"price"`
}
