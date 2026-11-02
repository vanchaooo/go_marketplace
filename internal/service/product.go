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

func (p *ProductService) AddNewProduct(ctx context.Context, name string, price, stock int64) (model.Product, error) {
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
	product := model.Product {
		Name: name,
		Price: price,
		Stock: stock,
	}
	result, err := p.products.Add(ctx, &product)
	if err != nil {
		return model.Product{}, err
	}
	return result, nil
}

func (p *ProductService) UpdateProduct(ctx context.Context, productID int64, name string, price, stock int64) (model.Product, error) {
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

	updProduct := model.Product {
		ID: productID,
		Name: name,
		Price: price,
		Stock: stock,
	}
	updatedProduct, err := p.products.Update(ctx, updProduct)
	if err != nil {
		return model.Product{}, err
	}
	return updatedProduct, nil
}

func (p *ProductService) DeleteProduct(ctx context.Context, productID int64) error {
	if productID <= 0 {
		return ErrProductNotFound
	}
	exist, err := p.products.Exists(ctx, productID)
	if err != nil {
		return err
	}
	if !exist {
		return ErrProductNotFound
	}

	ok := p.products.Delete(ctx, productID)
	if ok != nil {
		return ok
	}
	return nil
}

func (p *ProductService) GetProduct(ctx context.Context, productID int64) (model.Product, error) {
	if productID <= 0 {
		return model.Product{}, ErrProductNotFound
	}
	exist, err := p.products.Exists(ctx, productID)
	if err != nil {
		return model.Product{}, err
	}
	if !exist {
		return model.Product{}, ErrProductNotFound
	}

	product, err := p.products.Get(ctx, productID)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (p *ProductService) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	products, err := p.products.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}