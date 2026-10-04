package main

type User struct {
	ID int64
	Name string
}

type Wallet struct {
	UserID int64
	Balance int64
}

type Product struct {
	ID int64
	Name string
	Price int64
	Stock int64
}

type Cart struct {
	UserID int64
	Items map[int64]int64
}

type OrderItem struct {
	ProductID int64
	ProductName string
	Price int64
	Quantity int64
}

type Order struct {
	OrderID int64
	UserID int64
	Product []OrderItem
	Total int64
	Status string
}