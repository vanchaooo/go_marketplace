package repository

import (
	"strings"
	"slices"
	"github.com/vanchaooo/go-marketplace/internal/model"
)

type ProductRepository struct {
	products map[int64]*model.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: make(map[int64]*model.Product),
	}
}

func (p ProductRepository) Add(product *model.Product) {
	p.products[product.ID] = product
}

func (p ProductRepository) Get(productID int64) (*model.Product, bool) {
	product, ok := p.products[productID]
	return product, ok
}

func (p ProductRepository) Search(name string) []int64 {
	name = strings.ToLower(name)
	result := []int64{}

	for id, product := range p.products {
		if strings.Contains(strings.ToLower(product.Name), name) {
			result = append(result, id)
		}
	}
	slices.Sort(result)
	return result
}

func (p ProductRepository) Delete(productID int64) {
	delete(p.products, productID)
}

func (p ProductRepository) Exists(productID int64) bool {
	_, ok := p.products[productID]
	return ok
}
