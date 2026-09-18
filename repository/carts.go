package repository

import (
	"fmt"
	"errors"
)

type Item struct {
	Name string
	Price int
	Amount int
}

var Cart = map[int]*Item {}

func GetQuantity() int {
	count := 0
	for _, v := range Cart{
		count += v.Amount
	}
	return count
}

func GetAllCart() {
	fmt.Printf("Корзина %d:\n", GetQuantity())
	summ := 0
	for k, v := range Cart {
		fmt.Printf("%d. Товар: %q,  Цена: %d, Количество: %d.\n", k, v.Name, v.Price, v.Amount)
		summ += v.Price
	}
	fmt.Println()
	fmt.Printf("Итого: %d\n", summ)
}

func AddToCart(product string, amount int) bool {
	if product == "" {
		fmt.Println(errors.New("Название товара не может быть пустым."))
		return false
	}

	for _, v := range Catalog {
		if v.Name == product {
			addprod := &Item{
				Name: v.Name,
				Price: v.Price * amount,
				Amount: amount,
			}
			fmt.Printf("%q успешно добавлен в корзину.\n", product)
			newPosition := len(Cart)+1
			Cart[newPosition] = addprod
			return true
		}
	}

	fmt.Printf("Не удалось найти %q в каталоге товаров.", product)
	return false
}

func DeteleFromCart(product string) bool {
	for k, v := range Cart {
		if v.Name == product {
			fmt.Printf("%q успешно удален из корзины.\n", product)
			delete(Cart, k)
			return true
		}
	}

	fmt.Printf("Не удалось найти %q в корзине.\n", product)
	return false
}

func AddAmount(product string) bool {
	if product == "" {
		fmt.Println(errors.New("Название товара не может быть пустым."))
	}

	for _, v := range Cart {
		if v.Name == product {
			v.Amount++
			for _, check := range Catalog {
				if check.Name == product {
					v.Price += check.Price
				}
			}
			fmt.Println("Количество успешно увеличино на 1.")
			return true
		}
	}

	fmt.Printf("Не удалось найти %q в корзине.\n", product)
	return false
}

func LowerAmount(product string) bool {
	if product == "" {
		fmt.Println(errors.New("Название товара не может быть пустым."))
	}

	for _, v := range Cart {
		if v.Name == product {
			v.Amount--
			for _, check := range Catalog {
				if check.Name == product {
					v.Price -= check.Price
				}
			}
			fmt.Println("Количество успешно уменьшено на 1.")
			return true
		}
	}

	fmt.Printf("Не удалось найти %q в корзине.\n", product)
	return false
}