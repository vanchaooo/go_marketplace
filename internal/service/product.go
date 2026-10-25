package service

import (
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

func (p *ProductService) AddNewProduct(productID int64, name string, price, stock int64) bool {
	if p.products.Exists(productID) {
		return false
	}

	product, ok := model.NewProduct(productID, name, price, stock)
	if !ok {
		return false
	}

	p.products.Add(&product)
	return true
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