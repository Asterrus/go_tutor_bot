package product

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Service struct {
	products []*Product
}

func NewService() *Service {
	return &Service{
		products: []*Product{},
	}
}

func (s *Service) List() []*Product {
	return s.products
}

func (s *Service) Get(idx int) (*Product, error) {
	fmt.Printf("Service GET. idx: %d, len(products): %d", idx, len(s.products))
	if idx >= len(s.products) {
		return nil, fmt.Errorf("Product with idx: %d not found", idx)
	}
	product := s.products[idx]
	return product, nil
}

func (s *Service) LoadProducts(path string) error {

	file, err := os.ReadFile(path)
	if err != nil {
		log.Printf("No products.json file found in bot/ directory")
	}
	fmt.Printf("file, err: %s, %s", file, err)
	err = json.Unmarshal(file, &s.products)
	if err != nil {
		log.Println(err)
	}
	log.Println("Products: ", s.products)
	return err
}
