package service

import (
	"context"
	"errors"
	"sort"
	"github.com/vanchaooo/go-marketplace/internal/model"
	"github.com/vanchaooo/go-marketplace/internal/repository"
)

var (
	ErrNotFound = errors.New("по вашему запросу ничего не найдено")
	ErrInvalidID = errors.New("некорректный ID") 
	ErrInvalidName = errors.New("имя должно содержать от 1 до 200 символов")
	ErrInvalidPrice = errors.New("цена должна быть не менее 0 рублей")
	ErrInvalidStock = errors.New("количество должно быть не менее 0 штук")
	ErrProductNotFound = errors.New("продукт не найден")
	ErrUserNotFound = errors.New("пользователь не найден")
	ErrInvalidCredentials = errors.New("неверный email или пароль")
	ErrInvalidToken = errors.New("недействительный токен")
	ErrEmailTaken = errors.New("email уже занят")
	ErrInvalidEmail = errors.New("некорректный email")
	ErrInvalidPassword = errors.New("пароль должен быть от 8 до 72 символов")
	ErrWalletNotFound = errors.New("кошелек не найден")
	ErrInsufficientFunds = errors.New("недостаточно средств")
	ErrAmountNEgative = errors.New("сумма должна быть не менее 1 рубля")
)

func userCart(
	ctx context.Context,
	users *repository.UserRepository,
	carts *repository.CartRepository,
	userID int64,
) (*model.Cart, bool) {
	exist, err := users.Exists(ctx, userID)
	if err != nil {
		return &model.Cart{}, false
	}
	if !exist {
		return &model.Cart{}, false
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