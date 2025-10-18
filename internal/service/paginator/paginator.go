package paginator

import (
	"math"
)

type PaginatorInterface[T any] interface {
	GetPaginatedItems(items []T, page int) []T
	TotalPages(items []T) int
}

type Paginator[T any] struct {
	ItemsOnPage int
}

func NewPaginator[T any](itemsOnPage int) PaginatorInterface[T] {
	return &Paginator[T]{ItemsOnPage: itemsOnPage}
}

func (p *Paginator[T]) GetPaginatedItems(items []T, page int) []T {
	startIndex := (page - 1) * p.ItemsOnPage
	endIndex := min(p.ItemsOnPage*page, len(items))

	return items[startIndex:endIndex]
}

func (p *Paginator[T]) TotalPages(items []T) int {
	return int(math.Ceil((float64(len(items)) / float64(p.ItemsOnPage))))
}
