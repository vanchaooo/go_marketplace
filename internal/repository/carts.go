package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vanchaooo/go-marketplace/internal/model"
)

var ErrProductNotFound = errors.New("продукт не найден")
var ErrCartItemNotFound = errors.New("продукт не найден в корзине")

type CartRepository struct {
	pool *pgxpool.Pool
}

func NewCartRepository(pool *pgxpool.Pool) *CartRepository {
	return &CartRepository{pool: pool}
}

func (c *CartRepository) Add(ctx context.Context, userID, productID, quantity int64) (model.CartItem, error) {
	const query = `
		INSERT INTO cart_items (user_id, product_id, quantity)
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id, product_id)
		DO UPDATE SET quantity = cart_items.quantity + EXCLUDED.quantity
		RETURNING product_id, quantity
	`

	var cartItem model.CartItem
	if err := c.pool.QueryRow(ctx, query, userID, productID, quantity).Scan(
		&cartItem.ProductID,
		&cartItem.Quantity,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return model.CartItem{}, ErrProductNotFound
		}
		return model.CartItem{}, err
	}
	return cartItem, nil
}

func (c *CartRepository) SetQuantity(ctx context.Context, userID, productID, quantity int64) (model.CartItem, error) {
	const query = `
		UPDATE cart_items
		SET quantity = $3
		WHERE user_id = $1 AND product_id = $2
		RETURNING product_id, quantity
	`

	var cartItem model.CartItem
	if err := c.pool.QueryRow(ctx, query, userID, productID, quantity).Scan(
		&cartItem.ProductID,
		&cartItem.Quantity,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.CartItem{}, ErrCartItemNotFound
		}
		return model.CartItem{}, err
	}
	return cartItem, nil
}

func (c *CartRepository) Get(ctx context.Context, userID int64) (model.Cart, error) {
	const query = `
		SELECT p.id, p.name, p.price * c.quantity, c.quantity
		FROM cart_items c
		JOIN products p ON p.id = c.product_id
		WHERE c.user_id = $1
		ORDER BY p.id
	`

	rows, err := c.pool.Query(ctx, query, userID)
	if err != nil {
		return model.Cart{} , err
	}
	defer rows.Close()

	cart := model.Cart{UserID: userID, Items: []model.CartItem{}}
	for rows.Next() {
		var cartItem model.CartItem
		if err := rows.Scan(&cartItem.ProductID, &cartItem.Name, &cartItem.Price, &cartItem.Quantity); err != nil {
			return model.Cart{}, err
		}
		cart.Items = append(cart.Items, cartItem)
	}

	if err := rows.Err(); err!= nil {
		return model.Cart{}, err
	}
	return cart, nil
}

func (c *CartRepository) Delete(ctx context.Context, userID, productID int64) error {
	const query = `
		DELETE from cart_items
		WHERE user_id = $1 AND product_id = $2
	`

	response, err := c.pool.Exec(ctx, query, userID, productID)
	if err != nil {
		return err
	}
	if response.RowsAffected() == 0 {
		return ErrCartItemNotFound
	}
	return nil
}