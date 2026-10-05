package main

import (
	"fmt"
	"math"
	"sort"
)

func (m *Marketplace) Checkout(userID int64) (int64, bool) {
	cart, ok := m.userCart(userID)
	if !ok {
		return 0, false
	}
	wallet, exists := m.wallets[userID]
	if !exists {
		return 0, false
	}
	if cart.IsEmpty() {
		return 0, false
	}

	items := make([]OrderItem, 0, len(cart.Items))
	for productID, quantity := range cart.Items {
		product, exists := m.products[productID]
		if !exists {
			return 0, false
		}
		if !product.IsAvailable(quantity) {
			return 0, false
		}

		item, ok := NewOrderItem(*product, quantity)
		if !ok {
			return 0, false
		}
		items = append(items, item)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].ProductID < items[j].ProductID
	})

	total, ok := CalculateOrderTotal(items)
	if !ok {
		return 0, false
	}
	if !wallet.HasEnoughMoney(total) {
		return 0, false
	}

	orderID := m.nextOrderID
	if orderID <= 0 || orderID == math.MaxInt64 {
		return 0, false
	}
	_, taken := m.orders[orderID]
	if taken {
		return 0, false
	}

	order, ok := NewOrder(orderID, userID, items)
	if !ok {
		return 0, false
	}

	wallet.Withdraw(total)
	for _, item := range items {
		m.products[item.ProductID].Reserve(item.Quantity)
	}
	m.orders[orderID] = &order

	cart.Clear()
	line := fmt.Sprintf("Покупка: заказ %d на сумму %d", orderID, total)
	m.operationHistory[userID] = append(m.operationHistory[userID], line)
	m.nextOrderID++

	return orderID, true
}

func (m *Marketplace) CancelOrder(userID, orderID int64) bool {
	_, exists := m.users[userID]
	if !exists {
		return false
	}
	wallet, exists := m.wallets[userID]
	if !exists {
		return false
	}
	order, exists := m.orders[orderID]
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
		product, exists := m.products[item.ProductID]
		if !exists {
			return false
		}
		if product.Stock > math.MaxInt64-item.Quantity {
			return false
		}
	}

	for _, item := range order.Product {
		m.products[item.ProductID].AddStock(item.Quantity)
	}
	wallet.TopUp(order.Total)

	order.MarkCancelled()
	line := fmt.Sprintf("Отмена: заказ %d, возвращено %d", orderID, order.Total)
	m.operationHistory[userID] = append(m.operationHistory[userID], line)

	return true
}