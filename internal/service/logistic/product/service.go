package product

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

type ProductMap map[int]*Product

type ProductService interface {
	Get(productID int) (*Product, error)
	List() []*Product
	Create(title string, price float64) int
	Update(ProductID int, title string, price float64)
	Remove(ProductID int)
}

type Service struct {
	products ProductMap
}

func NewService() ProductService {
	service := Service{
		products: ProductMap{},
	}
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "data", "products.json")
	load_err := service.LoadProducts(path)

	if load_err != nil {
		log.Println("Unable to load products")
	}
	return &service
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

func (s *Service) Get(productID int) (*Product, error) {
	product, ok := s.products[productID]
	if !ok {
		return nil, fmt.Errorf("no product with ID %d", productID)
	}
	return product, nil
}

func (s *Service) Remove(productID int) {
	delete(s.products, productID)
}

func (s *Service) getNewID() int {
	maxKey := 0
	for key := range s.products {
		maxKey = max(maxKey, key)
	}
	return maxKey + 1
}

func (s *Service) Create(title string, price float64) int {
	productID := s.getNewID()
	newProduct := Product{
		ID:    productID,
		Title: title,
		Price: price,
	}
	s.products[productID] = &newProduct
	return productID
}
func (s *Service) Update(productID int, title string, price float64) {
	editedProduct, _ := s.Get(productID)
	editedProduct.Price = price
	editedProduct.Title = title
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
