package service

import (
	"context"
	"math"
	"strings"
	"unicode/utf8"

	"github.com/vanchaooo/go-marketplace/internal/model"
	"github.com/vanchaooo/go-marketplace/internal/repository"
)

type ProductService struct {
	products *repository.ProductRepository
}

func NewProductService(
	products *repository.ProductRepository,
) *ProductService {
	return &ProductService {
		products: products,
	}
}

func (p *ProductService) AddNewProduct(ctx context.Context, productID int64, name string, price, stock int64) (model.Product, error) {
	name = strings.TrimSpace(name)
	if length := utf8.RuneCountInString(name); length < 1 || length > 200 {
		return model.Product{}, ErrInvalidName
	}
	if price < 0 || price > math.MaxInt64 {
		return model.Product{}, ErrInvalidPrice
	}
	if stock < 0 || stock > math.MaxInt64 {
		return model.Product{}, ErrInvalidStock
	}
	model := &model.Product {
		ID: productID,
		Name: name,
		Price: price,
		Stock: stock,
	}
	p.products.Add(model)
	return *model, nil
}

func (p *ProductService) ChangeProductPrice(productID, newPrice int64) bool {
	product, ok := p.products.Get(productID)
	if !ok {
		return false
	}

	return product.SetPrice(newPrice)
}

func (p *ProductService) ChangeProductStock(productID, newStock int64) bool {
	product, ok := p.products.Get(productID)
	if !ok {
		return false
	}

	return product.SetStock(newStock)
}

func (p *ProductService) RenameProduct(productID int64, newName string) bool {
	product, ok := p.products.Get(productID)
	if !ok {
		return false
	}

	return product.Rename(newName)
}

func (p *ProductService) GetProduct(productID int64) (model.Product, bool) {
	product, ok := p.products.Get(productID)
	if !ok {
		return model.Product{}, false
	}

	return *product, true
}

func (p *ProductService) SearchProduct(name string) []int64 {
	return p.products.Search(name)
}