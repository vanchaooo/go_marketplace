package service

import (
	"context"
	"errors"

	"github.com/vanchaooo/go-marketplace/internal/model"
	"github.com/vanchaooo/go-marketplace/internal/repository"
)

type CartService struct {
	users *repository.UserRepository
	products *repository.ProductRepository
	carts *repository.CartRepository
}

func NewCartService(
	users *repository.UserRepository,
	products *repository.ProductRepository,
	carts *repository.CartRepository,
) *CartService {
	return &CartService{
		users: users,
		products: products,
		carts: carts,
	}
}

func (c *CartService) AddToCart(ctx context.Context, userID, productID, quantity int64) (model.CartItem, error) {
	if userID <= 0 {
		return model.CartItem{}, ErrInvalidID
	}
	if productID <= 0 {
		return model.CartItem{}, ErrInvalidID
	}
	if quantity <= 0 {
		return model.CartItem{}, ErrInvalidQuantity
	}
	userExist, err := c.users.Exists(ctx, userID)
	if err != nil {
		return model.CartItem{}, err
	}
	if !userExist {
		return model.CartItem{}, ErrUserNotFound
	}
	productExist, err := c.products.Exists(ctx, productID)
	if err != nil {
		return model.CartItem{}, err
	}
	if !productExist {
		return model.CartItem{}, ErrProductNotFound
	}
	
	cartItem, err := c.carts.Add(ctx, userID, productID, quantity)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			return model.CartItem{}, ErrProductNotFound
		}
		return model.CartItem{}, err
	}
	return cartItem, nil
}

func (c *CartService) ChangeQuantity(ctx context.Context, userID, productID, quantity int64) (model.CartItem, error) {
	if userID <= 0 || productID <= 0 {
		return model.CartItem{}, ErrInvalidID
	}
	if quantity <= 0 {
		return model.CartItem{}, ErrInvalidQuantity
	}

	cartItem, err := c.carts.SetQuantity(ctx, userID, productID, quantity)
	if err != nil {
		if errors.Is(err, repository.ErrCartItemNotFound) {
			return model.CartItem{}, ErrCartItemNotFound
		}
		return model.CartItem{}, err
	}
	return cartItem, nil
}

func (c *CartService) GetCart(ctx context.Context, userID int64) (model.Cart, error) {
	if userID <= 0 {
		return model.Cart{}, ErrInvalidID
	}
	userExist, err := c.users.Exists(ctx, userID)
	if err != nil {
		return model.Cart{}, err
	}
	if !userExist {
		return model.Cart{}, ErrUserNotFound
	}

	cart, err := c.carts.Get(ctx, userID)
	if err != nil {
		return model.Cart{}, err
	}
	return cart, nil
}

func (c *CartService) DeleteItem(ctx context.Context, userID, productID int64) error {
	if userID <= 0 {
		return ErrInvalidID
	}
	if productID <= 0 {
		return ErrInvalidID
	}
	userExist, err := c.users.Exists(ctx, userID)
	if err != nil {
		return err
	}
	if !userExist {
		return ErrUserNotFound
	}

	if err := c.carts.Delete(ctx, userID, productID); err != nil {
		if errors.Is(err, repository.ErrCartItemNotFound) {
			return ErrCartItemNotFound
		}
		return err
	}
	return nil
}