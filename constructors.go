package main

import (
	"errors"
	"fmt"
	"github.com/vanchaooo/go-marketplace/repository"
)

func NewUser(name string) (User, bool) {
	if name == "" {
		fmt.Println(errors.New("Введите свое имя"))
		return User{}, false
	}


	id := repository.GetAllUsers() + 1
	user := User{
		ID: id,
		Name: name, 
		Balance: 0,
	}
	return user, true
}

func NewProduct(name string, price int, stock int) (Product, bool) {
	if name == "" {
		fmt.Println(errors.New("Название товара не может быть пустым"))
		return Product{}, false
	}
	if price == 0 {
		fmt.Println(errors.New("Цена товара должна быть больше 0"))
		return Product{}, false
	}
	if stock == 0 {
		fmt.Println(errors.New("Для добавления товара, он должен быть в наличии"))
		return Product{}, false
	}

	id := repository.GetCatalog() + 1
	product := Product{
		ID: id,
		Name: name,
		Price: price,
		Stock: stock,
	}

	return product, true
}

func NewCart(userID int) (Cart, bool) {
	if repository.UserExists(userID) {
		cart := Cart{
			UserID: userID,
			Items: make(map[int]int),
		}
		return cart, true
	}
	return Cart{}, false
}

func NewOrderItem(product Product, quantity int,) (OrderItem, bool) {
	if quantity == 0 {
		return OrderItem{}, false
	}

	item := OrderItem {
		ProductID: product.ID,
		ProductName: product.Name,
		Price: product.Price * quantity,
		Quantity: quantity,
	}
	return item, true
}

func NewOrder(id int, userID int, productID int) (Order, bool) {
	if repository.UserExists(userID) {
		prod, ok := repository.GetProduct(productID)

		if ok {
			order := Order{
				OrderID: len(repository.GetOrders(userID))+1,
				UserID: userID,
				ProductID: prod.ID,
				Total: repository.ProductCost(userID, prod.ID),
				Status: "Оплачен",
			}
			return order, true
		}
	}
	return Order{}, false
}

func CalculateOrderTotal(items []OrderItem) int {
	total := 0
	for _, item := range items {
		total += item.Price * item.Quantity
	}
	return total
}
 
func CopyOrderItems(items []OrderItem) []OrderItem {
	itemsCopy := make([]OrderItem, len(items))
	copy(itemsCopy, items)
	return itemsCopy
}