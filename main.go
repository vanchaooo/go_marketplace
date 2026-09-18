package main

import (
	"fmt"
	"github.com/vanchaooo/go-marketplace/repository"
)

func main() {
	repository.AddToCart("Кожаный кошелек", 10)
	repository.AddToCart("Ноутбук для работы 15.6", 1)
	fmt.Println()
	repository.GetAllCart()
	
}