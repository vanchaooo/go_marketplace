package repository

import (
	"errors"
	"fmt"
	"slices"
)

type Item struct {
	ProductID int
	Amount int
}

var Cart = map[int][]*Item {}

func TotalCartCost(userID int) int {
	if UserExists(userID) {
		sum := 0
		for _, v := range Cart[userID] {
			prod, _ := GetProduct(v.ProductID)
			sum += prod.Price * v.Amount
			return sum
		}
	}

	return 0
}

func ProductCost(userID int, productID int) int {
	prod, _ := GetProduct(productID)
	if UserExists(userID) {
		if productID == 0 {
			fmt.Println(errors.New("ID продукта не может быть равно 0."))
			return 0
		}
		for _, v := range Cart[userID] {
			if v.ProductID == prod.ID {
				result := prod.Price * v.Amount
				return result
			}
		}
	}

	fmt.Println(errors.New("Не удалось простичать сумму товара."))
	return 0
}

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
	if UserExists(userID) {
		total := 0
		summ := 0

		fmt.Printf("Корзина пользователя %d:\n", userID)
		for _, v := range Cart[userID] {
			product, ok := GetProduct(v.ProductID)
			if ok {
				fmt.Printf("Товар %q, Цена: %d, Количество: %d.\n", product.Name, product.Price, v.Amount)
				total += v.Amount
				summ += product.Price * v.Amount
			} else {
				fmt.Printf("Товар с ID %d не найден.\n", v.ProductID)
			}
		}
		fmt.Println()
		fmt.Printf("Товаров - %d.\nОбщая сумма товаров - %d.\n", total, summ)
		return true
	}

	fmt.Printf("Не удалось найти пользователя с таким ID.")
	return false
}

func AddToCart(userID int, productID int) bool {
	prod, ok := GetProduct(productID)
	if UserExists(userID) {
		if productID == 0 {
			fmt.Println("ID товара не может быть равно 0.")
			return false
		}
		if ok {
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
			fmt.Printf("%q добавлен в корзину.\n", prod.Name)
			Cart[userID] = append(Cart[userID], addprod)
			return true
		}
	}

	fmt.Printf("Не удалось найти %q в каталоге товаров.\n", prod.Name)
	return false
}

func DeleteFromCart(userID int, productID int) bool {
	prod, _ := GetProduct(productID)
	if UserExists(userID) {
		for k, _ := range Cart[userID] {
			if productID == prod.ID {
				fmt.Printf("%q удален из корзины.\n", prod.Name)
				Cart[userID] = slices.Delete(Cart[userID], k, k+1)
				return true
			}
		}
	}

	fmt.Printf("Не удалось найти %q в корзине.\n", prod.Name)
	return false
}

func AddAmount(userID int, productID int) bool {
	prod, _ := GetProduct(productID)
	if UserExists(userID) {
		if productID == 0 {
			fmt.Println("ID продукта не может быть равно 0")
			return false
		}

		for _, v := range Cart[userID] {
			if prod.ID == v.ProductID {
				v.Amount++
				fmt.Printf("%q увеличен на 1 шт.\n", prod.Name)
				return true
			}
		}
	}
	fmt.Printf("Не удалось найти %q в корзине.\n", prod.Name)
	return false
}

func LowerAmount(userID int, productID int) bool {
	prod, _ := GetProduct(productID)
	if UserExists(userID) {
		if productID == 0 {
			fmt.Println("ID продукта не может быть равно 0")
			return false
		}

		for k, v := range Cart[userID] {
			if prod.ID == v.ProductID {
				if v.Amount > 1 {
					v.Amount--
					fmt.Printf("%q увеличен на 1 шт.\n", prod.Name)
					return true
				} else {
					Cart[userID] = slices.Delete(Cart[userID], k, k+1)
					fmt.Printf("%q удален из корзины.\n", prod.Name)
					return true
				}
			}
		}
	}

	fmt.Printf("Не удалось найти %q в корзине.\n", prod.Name)
	return false
}