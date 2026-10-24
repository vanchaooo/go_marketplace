package repository

import (
	"github.com/vanchaooo/go-marketplace/internal/model"
)

type CartRepository struct {
	carts map[int64]*model.Cart
}

func NewCartRepository() *CartRepository {
	return &CartRepository{
		carts: make(map[int64]*model.Cart),
	}
}

func (c CartRepository) Add(cart *model.Cart) {
	c.carts[cart.UserID] = cart
}

func (c CartRepository) Get(userID int64) (*model.Cart, bool) {
	cart, ok := c.carts[userID]
	return cart, ok
}

func (c CartRepository) Delete(userID int64) {
	delete(c.carts, userID)
}