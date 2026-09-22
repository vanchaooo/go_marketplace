package main

import (
	"fmt"
	"github.com/vanchaooo/go-marketplace/service"
)

func main() {
	fmt.Println("======================================")
	fmt.Println()
	
	service.AddToCart(9, "Ноутбук для работы 15.6")
	service.AddAmount(9, "Ноутбук для работы 15.6")

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println()

	service.CheckCart(9)

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println()

	service.UpdateBalance(9, 200000)

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println()

	service.Order(9, "Ноутбук для работы 15.6")

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println()

	service.CheckOrdersHistory(9)

	fmt.Println()
	fmt.Println("======================================")
}