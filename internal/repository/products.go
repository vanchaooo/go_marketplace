package repository

import (
	"context"
	"slices"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vanchaooo/go-marketplace/internal/model"
)

type ProductRepository struct {
	pool *pgxpool.Pool
}

func NewProductRepository(pool *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{
		pool: pool,
	}
}

func (p ProductRepository) Add(ctx context.Context, newProduct *model.Product) (model.Product, error) {
	const query = `
		INSERT INTO products
		VALUES($1, $2, $3)
		RETURNING id, name, price, stock
	`
	var product model.Product
	if err := p.pool.QueryRow(ctx, query, product.Name, product.Price, product.Stock).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Stock,
	); err != nil {
		return model.Product{}, err
	}
	return product, nil
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
