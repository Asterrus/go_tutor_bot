package paginator

import "math"

type Paginator[T any] struct {
	ItemsOnPage int
}

func NewPaginator[T any](itemsOnPage int) *Paginator[T] {
	return &Paginator[T]{ItemsOnPage: itemsOnPage}
}

func (p *Paginator[T]) GetPaginatedItems(items []T, page int) []T {
	return items[(page-1)*p.ItemsOnPage : p.ItemsOnPage*page]
}

func (p *Paginator[T]) TotalPages(items []T) int {
	return int(math.Ceil(float64((len(items) / p.ItemsOnPage))))
}
