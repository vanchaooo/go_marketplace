package main

import (
	"unicode/utf8"
	"strings"
	"math"
)

func normalizeName(name string) (string, bool) {
	name = strings.TrimSpace(name)
 
	length := utf8.RuneCountInString(name)
	if length < 1 || length > 200 {
		return "", false
	}
	return name, true
}

func NewUser(userID int64, name string) (User, bool) {
	name, ok := normalizeName(name)
	if ok {
		return User{ID: userID, Name: name}, true
	}

	return User{}, false
}

func NewWallet(userID int64) (Wallet, bool) {
	if userID <= 0 {
		return Wallet{}, false
	}
	
	return Wallet{UserID: userID, Balance: 0}, true
}

func NewProduct(productID int64, name string, price int64, stock int64) (Product, bool) {
	if productID <= 0 {
		return Product{}, false
	}

	name, ok := normalizeName(name)
	if !ok {
		return Product{}, false
	}
	if price <= 0 {
		return Product{}, false
	}
	if stock < 0 {
		return Product{}, false
	}

	return Product{ID: productID, Name: name, Price: price, Stock: stock}, true
}

func NewCart(userID int64) (Cart, bool) {
	if userID <= 0 {
		return Cart{}, false
	}
	
	return Cart{UserID: userID, Items: make(map[int64]int64)}, true
}
 
func isValidOrderItem(product OrderItem) bool {
	return product.ProductID > 0 && product.Price > 0 && product.Quantity > 0 && product.ProductName != ""
}
 
func NewOrderItem(product Product, quantity int64) (OrderItem, bool) {
	item := OrderItem{
		ProductID:   product.ID,
		ProductName: product.Name,
		Price:       product.Price,
		Quantity:    quantity,
	}
 
	if isValidOrderItem(item) {
		return item, true
	}

	return OrderItem{}, false
}

func CalculateOrderTotal(products []OrderItem) (int64, bool) {
	var total int64 = 0
 
	for _, item := range products {
		if !isValidOrderItem(item) {
			return 0, false
		}

		if item.Price > math.MaxInt64/item.Quantity {
			return 0, false
		}
		line := item.Price * item.Quantity
		if total > math.MaxInt64-line {
			return 0, false
		}
		total += line
	}
 
	return total, true
}

func CopyOrderItems(products []OrderItem) []OrderItem {
	itemsCopy := make([]OrderItem, len(products))
	copy(itemsCopy, products)
	return itemsCopy
}

func NewOrder(orderID int64, userID int64, items []OrderItem) (Order, bool) {
	if orderID <= 0 || userID <= 0 {
		return Order{}, false
	}
	if len(items) == 0 {
		return Order{}, false
	}

 	itemsCopy := CopyOrderItems(items)
 
	total, ok := CalculateOrderTotal(itemsCopy)
	if !ok {
		return Order{}, false
	}
 
	order := Order{
		OrderID: orderID,
		UserID: userID,
		Product: itemsCopy,
		Total: total,
		Status: "paid",
	}
	return order, true
}