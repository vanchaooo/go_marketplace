package main

type User struct {
	ID int
	Name string
	Balance int
}

type Product struct {
	ID int
	Name string
	Price int
	Stock int
}

type Cart struct {
	UserID int
	Items map[int]int
}

type OrderItem struct {
	ProductID int
	ProductName string
	Price int
	Quantity int
}

type Order struct {
	OrderID int
	UserID int
	ProductID int
	Total int
	Status string
}