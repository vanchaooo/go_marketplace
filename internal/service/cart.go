package service

// import (
// 	"context"

// 	"github.com/vanchaooo/go-marketplace/internal/model"
// 	"github.com/vanchaooo/go-marketplace/internal/repository"
// )

// type CartService struct {
// 	users    *repository.UserRepository
// 	carts    *repository.CartRepository
// 	products *repository.ProductRepository
// }

// func NewCartService(
// 	users *repository.UserRepository,
// 	carts *repository.CartRepository,
// 	products *repository.ProductRepository,
// ) *CartService {
// 	return &CartService{users: users, carts: carts, products: products}
// }

// func (c *CartService) AddToCart(ctx context.Context, userID, productID, quantity int64) bool {
// 	cart, ok := userCart(c.users, c.carts, userID)
// 	if !ok {
// 		return false
// 	}
// 	exist, err :=c.products.Exists(ctx, productID)
// 	if err != nil {
// 		return false
// 	}
// 	if !exist {
// 		return false
// 	}

// 	return cart.Add(productID, quantity)
// }

// func (c *CartService) GetCart(userID int64) (model.Cart, bool) {
// 	cart, ok := userCart(c.users, c.carts, userID)
// 	if !ok {
// 		return model.Cart{}, false
// 	}

// 	return cart.Copy(), true
// }

// func (c *CartService) UpdateQuantity(ctx context.Context, userID, productID, quantiry int64) bool {
// 	cart, ok := userCart(c.users, c.carts, userID)
// 	if !ok {
// 		return false
// 	}
// 	exist, err :=c.products.Exists(ctx, productID)
// 	if err != nil {
// 		return false
// 	}
// 	if !exist {
// 		return false
// 	}
// 	if _, ok := cart.Items[productID]; !ok {
// 		return false
// 	}
	
// 	return cart.SetQuantity(productID, quantiry)
// }

// func (c *CartService) DeleteFromCart(userID, productID int64) bool {
// 	cart, ok := userCart(c.users, c.carts, userID)
// 	if !ok {
// 		return false
// 	}
	
// 	return cart.Remove(productID)
// }

// func (c *CartService) ClearCart(userID int64) bool {
// 	cart, ok := userCart(c.users, c.carts, userID)
// 	if !ok {
// 		return false
// 	}

// 	cart.Clear()
// 	return true
// }

// func (c *CartService) CalculateCartTotal(ctx context.Context, userID int64) (int64, bool) {
// 	cart, ok := userCart(c.users, c.carts, userID)
// 	if !ok {
// 		return 0, false
// 	}
	
// 	items, ok := buildOrderItem(ctx, cart, c.products, false)
// 	if !ok {
// 		return 0, false
// 	}

// 	return model.CalculateOrderTotal(items)
// }