package service

import (
	"context"
	"errors"
	"sort"

	"github.com/vanchaooo/go-marketplace/internal/model"
	"github.com/vanchaooo/go-marketplace/internal/repository"
)

var ErrNotFound = errors.New("По вашему запросу ничего не найдено.")
var ErrInvalidName = errors.New("Имя должно содержать от 1 до 200 символов.")
var ErrInvalidPrice = errors.New("Цена должна быть не менее 0 рублей.")
var ErrInvalidStock = errors.New("Количество должно быть не менее 0 штук.")
var ErrProductNotFound = errors.New("Продукт не был найден.")


func userCart(
	users *repository.UserRepository,
	carts *repository.CartRepository,
	userID int64,
) (*model.Cart, bool) {
	if !users.Exists(userID) {
		return nil, false
	}
 
	return carts.Get(userID)
}

func buildOrderItem(
	ctx context.Context,
	cart *model.Cart,
	products *repository.ProductRepository,
	stock bool,
) ([]model.OrderItem, bool) {
	items := make([]model.OrderItem, 0, len(cart.Items))

	for productID, quantity := range cart.Items {
		product, err := products.Get(ctx, productID)
		if err != nil {
			return nil, false
		}
		if stock && !product.IsAvailable(quantity) {
			return nil, false
		}

		item, ok := model.NewOrderItem(product, quantity)
		if !ok {
			return nil, false
		}
		items = append(items, item)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].ProductID < items[j].ProductID
	})
	return items, true
}