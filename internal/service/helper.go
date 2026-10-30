package service

import (
	"errors"
	"sort"
	"github.com/vanchaooo/go-marketplace/internal/model"
	"github.com/vanchaooo/go-marketplace/internal/repository"
)

var ErrInvalidName = errors.New("Имя должно содержать от 1 до 200 символов.")
var ErrInvalidPrice = errors.New("Цена должна быть не менее 0 рублей.")
var ErrInvalidStock = errors.New("Количество должно быть не менее 0 штук.")


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
	cart *model.Cart,
	products *repository.ProductRepository,
	stock bool,
) ([]model.OrderItem, bool) {
	items := make([]model.OrderItem, 0, len(cart.Items))

	for productID, quantity := range cart.Items {
		product, ok := products.Get(productID)
		if !ok {
			return nil, false
		}
		if stock && !product.IsAvailable(quantity) {
			return nil, false
		}

		item, ok := model.NewOrderItem(*product, quantity)
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