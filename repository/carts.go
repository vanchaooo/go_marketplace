package repository

import (
	"errors"
	"fmt"
	"slices"

	"github.com/vanchaooo/go-marketplace/repository"
)

type Item struct {
	ProductID int
	Amount int
}

var Cart = map[int][]*Item {}

func GetQuantity() int {
	count := 0
	for _, v := range Cart {
		for _, j := range v {
			count += j.Amount
		}
	}
	return count
}

func GetCart(userID int) bool {
	if userID == 0 {
		fmt.Println(errors.New("Введите ваш ID."))
		return false
	}

	total := 0
	summ := 0
	fmt.Printf("Корзина пользователя %d:\n", userID)
	for _, v := range Cart[userID] {
		fmt.Printf("Товар %q, Цена: %d, Количество: %d.\n", v.Name, v.Price, v.Amount)
		total += v.Amount
		summ += v.Price
	}
	fmt.Println()
	fmt.Printf("Товаров - %d.\nОбщая сумма товаров - %d.\n", total, summ)
	return true
}

func AddToCart(userID int, productID int) bool {
	if userID == 0 {
		fmt.Println(errors.New("Введите ваш ID."))
		return false
	}
	product := repository.GetProduct(productID)
	if !product {
		return false
	}


	for _, v := range Cart[userID] {
		if v.ProductID == productID {
			AddAmount(userID, productID)
			return true
		}
	}

	addprod := &Item {
		ProductID: productID,
		Amount: 1,
	}
	fmt.Printf("%q добавлен в корзину.\n", product)
	Cart[userID] = append(Cart[userID], addprod)
	return true
}

// ===================================== Переделать все по ID =================================================== 

func DeleteFromCart(userID int, product string) bool {
	for k, j := range Cart[userID] {
		if j.Name == product {
			fmt.Printf("%q удален из корзины.\n", product)
			Cart[userID] = slices.Delete(Cart[userID], k, k+1)
			return true
		}
	}

	fmt.Printf("Не удалось найти %q в корзине.\n", product)
	return false
}

func AddAmount(userID int, productID int) bool {
	if product == "" {
		fmt.Println(errors.New("Название товара не может быть пустым."))
	}

	for _, j := range Cart[userID] {
		if j.Name == product {
			j.Amount++
			for _, check := range Catalog {
				if check.Name == product {
					j.Price += check.Price
				}
			}
			fmt.Printf("%q увеличен на 1.\n", product)
			return true
		}
	}

	fmt.Printf("Не удалось найти %q в корзине.\n", product)
	return false
}

func LowerAmount(userID int, product string) bool {
	if product == "" {
		fmt.Println(errors.New("Название товара не может быть пустым."))
	}

	for k, j := range Cart[userID] {
		if j.Name == product {
			if j.Amount > 1 {
				j.Amount--
				for _, check := range Catalog {
					if check.Name == product {
						j.Price -= check.Price
					}
				}
				fmt.Printf("%q уменьшен на 1.\n", product)
				return true
			} else {
				Cart[userID] = slices.Delete(Cart[userID], k, k+1)
				fmt.Printf("%q удален из корзины.\n", product)
				return true
			}
		}
	}

	fmt.Printf("Не удалось найти %q в корзине.\n", product)
	return false
}