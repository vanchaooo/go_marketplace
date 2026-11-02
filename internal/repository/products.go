package repository

import (
	"errors"
	"context"
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
		INSERT INTO products (name, price, stock)
		VALUES ($1, $2, $3)
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

func (p ProductRepository) Update(ctx context.Context, product model.Product) (model.Product, error) {
	const query = `
		UPDATE products
		SET name = $2, price = $3, stock = $4
		WHERE id = $1
	`
	var updatedProduct model.Product
	if err := p.pool.QueryRow(ctx, query, product.ID, product.Name, product.Price, product.Stock).Scan(
		&updatedProduct.ID,
        &updatedProduct.Name,
        &updatedProduct.Price,
        &updatedProduct.Stock,
	); err != nil {
		return model.Product{}, err
	}
	
	return updatedProduct, nil
}

func (p ProductRepository) Get(ctx context.Context, productID int64) (model.Product, error) {
	const query = `
		SELECT id, name, price, stock
		FROM products
		WHERE id = $1
	`
	var product model.Product
	if err := p.pool.QueryRow(ctx, query, productID).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Stock,
	); err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (p ProductRepository) GetAll(ctx context.Context) ([]model.Product, error) {
	const query = `
		SELECT *
		FROM products
	`
	var productsList []model.Product

	rows, err := p.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock); err != nil {
			return nil, err
		}
		productsList = append(productsList, product)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return productsList, nil
}

func (p ProductRepository) Delete(ctx context.Context, productID int64) error {
	const query = `
		DELETE FROM products
		WHERE id = $1
	`
	response, err := p.pool.Exec(ctx, query, productID)
	if err != nil {
		return err
	}
	if response.RowsAffected() == 0 {
		return errors.New("Продукт не был найден.")
	}
	return nil
}

func (p ProductRepository) Exists(ctx context.Context, productID int64) (bool, error) {
	const query = `
		SELECT EXISTS (
			SELECT id
			FROM products
			WHERE id = $1
		)
	`
	var exist bool
	if err := p.pool.QueryRow(ctx, query, productID).Scan(&exist); err != nil {
		return false, err
	}
	return true, nil
}
