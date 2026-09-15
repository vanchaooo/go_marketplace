package main

import (
	"fmt"
	// "github.com/vanchaooo/go-marketplace/users"
	// "github.com/vanchaooo/go-marketplace/balance"
	"github.com/vanchaooo/go-marketplace/catalog"
	"github.com/vanchaooo/go-marketplace/cart"
	// "github.com/vanchaooo/go-marketplace/orders"
)

func main() {
	// users.AddUser("Mark")
	// users.GetAllusers()

	// fmt.Println()
	// fmt.Println("============================================================")
	// fmt.Println()

	// balance.TopUpBalance(users.Users_list, 5, 10079)
	// balance.GetEveryoneBalance()

	// fmt.Println()
	// fmt.Println("============================================================")
	// fmt.Println()

	catalog.AddNewSection("Столы")
	catalog.AddNewPosition("Столы", "Стол с механизмом подъема", 45000)
	catalog.GetCatalog()

	fmt.Println()
	fmt.Println("============================================================")
	fmt.Println()

	cart.AddToCart("Аудио", "AirPods Pro 2", 1)
	cart.AddToCart("Телефоны", "iPhone 15", 2)
	cart.AddToCart("Умный дом", "Яндекс Станция Макс", 4)
	fmt.Println()
	cart.GetCart()
	fmt.Println()
	cart.ReduceQuantity("AirPods Pro 2")
	cart.ReduceQuantity("Яндекс Станция Макс")
	cart.AddQuantity("iPhone 15")
	fmt.Println()
	cart.GetCart()

	// fmt.Println()
	// fmt.Println("============================================================")
	// fmt.Println()

	// orders.Order("AirPods Pro 2")
	// orders.Order("iPhone 15")
	// orders.GetAllOrders()
	// orders.CancelOrder("iPhone 15")
	// orders.GetAllOrders()
	// orders.OrdersHistory()
}