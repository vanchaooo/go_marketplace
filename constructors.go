package main

import (
	"unicode/utf8"
	"strings"
	"math"
	"github.com/vanchaooo/go-marketplace/repository"
)

func normalizeName(name string) (string, bool) {
	name = strings.TrimSpace(name)
 
	length := utf8.RuneCountInString(name)
	if length < 1 || length > 200 {
		return "", false
	}
	return name, true
}

func NewUser(name string) (User, bool) {
	name, ok := normalizeName(name)
	if ok {
		return User{ID: repository.GetAllUsers()+1, Name: name}, true
	}

	return User{}, false
}

func NewWallet(userID int64) (Wallet, bool) {
	if repository.UserExists(userID) {
		return Wallet{UserID: userID, Balance: 0}, true
	}
	
	return Wallet{}, false
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
	if repository.UserExists(userID) {
		return Cart{UserID: userID, Items: make(map[int64]int64)}, true
	}
	
	return Cart{}, false
}
 
func isValidOrderItem(item OrderItem) bool {
	return item.ProductID > 0 && item.Price > 0 && item.Quantity > 0 && item.ProductName != ""
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

func CalculateOrderTotal(items []OrderItem) (int64, bool) {
	var total int64 = 0
 
	for _, item := range items {
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

func CopyOrderItems(items []OrderItem) []OrderItem {
	itemsCopy := make([]OrderItem, len(items))
	copy(itemsCopy, items)
	return itemsCopy
}

func NewOrder(id int64, userID int64, items []OrderItem) (Order, bool) {
	if id <= 0 || userID <= 0 {
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
		OrderID:     id,
		UserID: userID,
		Product:  itemsCopy,
		Total:  total,
		Status: "paid",
	}
	return order, true
}