package product

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

type ProductMap map[int]*Product

type Service struct {
	products ProductMap
}

func NewService() *Service {
	return &Service{
		products: ProductMap{},
	}
}

func (s *Service) List() []*Product {
	productSlice := []*Product{}
	for _, value := range s.products {
		productSlice = append(productSlice, value)
	}
	sort.Slice(
		productSlice,
		func(i, j int) (less bool) { return productSlice[i].ID < productSlice[j].ID },
	)
	return productSlice
}

func (s *Service) Get(idx int) (*Product, error) {
	product, ok := s.products[idx]
	if !ok {
		return nil, fmt.Errorf("no product with ID %d", idx)
	}
	return product, nil
}

func (s *Service) Delete(idx int) {
	delete(s.products, idx)
}

func (s *Service) getNewID() int {
	maxKey := 0
	for key := range s.products {
		maxKey = max(maxKey, key)
	}
	return maxKey + 1
}

func (s *Service) New(name string, price float64) int {
	productID := s.getNewID()
	newProduct := Product{
		ID:    productID,
		Title: name,
		Price: price,
	}
	s.products[productID] = &newProduct
	return productID
}
func (s *Service) LoadProducts(path string) error {

	file, err := os.ReadFile(path)
	if err != nil {
		log.Printf("No products.json file found in bot/ directory")
	}
	fmt.Printf("file, err: %s, %s", file, err)
	productMap := map[int]Product{}
	err = json.Unmarshal(file, &productMap)
	if err != nil {
		log.Println(err)
	}
	for key, value := range productMap {
		s.products[key] = &value
	}

	log.Println("Products: ", s.products)
	return err
}
