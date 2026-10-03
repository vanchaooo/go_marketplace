package main

import (
	"fmt"

	// "github.com/vanchaooo/go-marketplace/repository"
	"github.com/vanchaooo/go-marketplace/service"
)

func main() {
	service.AddToCart(9, 27)
	service.AddToCart(9, 20)
	service.AddToCart(9, 9)

	fmt.Println()

	service.AddAmount(9, 9)

	fmt.Println()

	service.CheckCart(9)

	fmt.Println()

	service.UpdateBalance(9, 200000)

	fmt.Println()

	service.Order(9, 9)

	fmt.Println()

	service.GetMyOrders(9)
}