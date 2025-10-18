package paginator

import (
	"fmt"
	"math"
)

type Paginator[T any] struct {
	ItemsOnPage int
}

func NewPaginator[T any](itemsOnPage int) *Paginator[T] {
	return &Paginator[T]{ItemsOnPage: itemsOnPage}
}

func (p *Paginator[T]) GetPaginatedItems(items []T, page int) []T {
	startIndex := (page - 1) * p.ItemsOnPage
	var endIndex int
	if len(items) >= p.ItemsOnPage*page {
		endIndex = p.ItemsOnPage * page
	} else {
		endIndex = len(items)
	}

	return items[startIndex:endIndex]
}

func (p *Paginator[T]) TotalPages(items []T) int {
	fmt.Println("p.ItemsOnPage", p.ItemsOnPage)
	fmt.Println("len(items)", len(items))
	fmt.Println("len(items) / p.ItemsOnPage", len(items)/p.ItemsOnPage)
	fmt.Println("float64((len(items) / p.ItemsOnPage)", float64((len(items) / p.ItemsOnPage)))
	fmt.Println("math.Ceil(float64((len(items) / p.ItemsOnPage)))", math.Ceil(float64((len(items) / p.ItemsOnPage))))
	return int(math.Ceil((float64(len(items)) / float64(p.ItemsOnPage))))
}
