package service

import (
	"fmt"
	"math"
	"github.com/vanchaooo/go-marketplace/internal/model"
	"github.com/vanchaooo/go-marketplace/internal/repository"
)

type OrderService struct {
	users    *repository.UserRepository
	wallets  *repository.WalletRepository
	products *repository.ProductRepository
	carts    *repository.CartRepository
	orders   *repository.OrderRepository
	history  *repository.HistoryRepository
}
 
func NewOrderService(
	users *repository.UserRepository,
	wallets *repository.WalletRepository,
	products *repository.ProductRepository,
	carts *repository.CartRepository,
	orders *repository.OrderRepository,
	history *repository.HistoryRepository,
) *OrderService {
	return &OrderService{
		users:    users,
		wallets:  wallets,
		products: products,
		carts:    carts,
		orders:   orders,
		history:  history,
	}
}

func (o *OrderService) MakeOrder(userID int64) (int64, bool) {
	cart, ok := userCart(o.users, o.carts, userID)
	if !ok {
		return 0, false
	}

	wallet, ok := o.wallets.Get(userID)
	if !ok {
		return 0, false
	}

	if cart.IsEmpty() {
		return 0, false
	}

	items, ok := buildOrderItem(cart, o.products, true)
	if !ok {
		return 0, false
	}

	total, ok := model.CalculateOrderTotal(items)
	if !ok {
		return 0, false
	}
	if !wallet.HasEnoughMoney(total) {
		return 0, false
	}
 
	order, ok := model.NewOrder(userID, items)
	if !ok {
		return 0, false
	}

	orderID, ok := o.orders.Add(&order)
	if !ok {
		return 0, false
	}

	wallet.Withdraw(total)
	for _, item := range items {
		product, _ := o.products.Get(item.ProductID)
		product.Reserve(item.Quantity)
	}
	cart.Clear()

	text := fmt.Sprintf("Покупка: Заказ №%d на сумму %d.\n", orderID, total)
	o.history.Add(userID, text)
	return orderID, true
}

func (s *OrderService) CancelOrder(userID, orderID int64) bool {
	if !s.users.Exists(userID) {
		return false
	}
	wallet, exists := s.wallets.Get(userID)
	if !exists {
		return false
	}
	order, exists := s.orders.Get(orderID)
	if !exists {
		return false
	}
	if order.UserID != userID {
		return false
	}
	if !order.IsPaid() {
		return false
	}
 
	if wallet.Balance > math.MaxInt64-order.Total {
		return false
	}
	for _, item := range order.Product {
		product, exists := s.products.Get(item.ProductID)
		if !exists {
			return false
		}
		if product.Stock > math.MaxInt64-item.Quantity {
			return false
		}
	}
 
	for _, item := range order.Product {
		product, _ := s.products.Get(item.ProductID)
		product.AddStock(item.Quantity)
	}
	wallet.TopUp(order.Total)
	order.MarkCancelled()
 
	line := fmt.Sprintf("Отмена: заказ %d, возвращено %d", orderID, order.Total)
	s.history.Add(userID, line)
 
	return true
}

func (s *OrderService) GetOrder(userID, orderID int64) (model.Order, bool) {
	if !s.users.Exists(userID) {
		return model.Order{}, false
	}
 
	order, exists := s.orders.Get(orderID)
	if !exists || order.UserID != userID {
		return model.Order{}, false
	}
 
	return order.Copy(), true
}
 
func (s *OrderService) GetOrderItems(userID, orderID int64) ([]model.OrderItem, bool) {
	order, ok := s.GetOrder(userID, orderID)
	if !ok {
		return nil, false
	}
 
	return order.Product, true
}
 
func (s *OrderService) GetUserOrders(userID int64) []int64 {
	result := []int64{}
 
	if !s.users.Exists(userID) {
		return result
	}
 
	for _, order := range s.orders.GetAll(userID) {
		result = append(result, order.OrderID)
	}
	return result
}
 
func (s *OrderService) GetUserHistory(userID int64) []string {
	return s.history.Get(userID)
}
 