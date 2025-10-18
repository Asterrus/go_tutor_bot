package paginator

import (
	"reflect"
	"testing"
)

func TestTotalPages(t *testing.T) {
	tests := []struct {
		name        string
		items       []int
		itemsOnPage int
		expected    int
	}{
		{"no items", []int{}, 2, 0},
		{"one full page", []int{1, 2}, 2, 1},
		{"two page", []int{1, 2, 3, 4}, 2, 2},
		{"two page, one item in second page", []int{1, 2, 3}, 2, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Paginator[int]{ItemsOnPage: tt.itemsOnPage}
			got := p.TotalPages(tt.items)
			if got != tt.expected {
				t.Fatalf("expected %d, got %d", tt.expected, got)
			}
		})
	}
}

func TestGetPaginatedItems(t *testing.T) {
	tests := []struct {
		name        string
		items       []int
		itemsOnPage int
		page        int
		expected    []int
	}{
		{"no items", []int{}, 2, 1, []int{}},
		{"one full page", []int{1, 2}, 2, 1, []int{1, 2}},
		{"two page", []int{1, 2, 3, 4}, 2, 2, []int{3, 4}},
		{"two page, one item in second page", []int{1, 2, 3}, 2, 2, []int{3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Paginator[int]{ItemsOnPage: tt.itemsOnPage}
			got := p.GetPaginatedItems(tt.items, tt.page)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Fatalf("expected %d, got %d", tt.expected, got)
			}
		})
	}
}
