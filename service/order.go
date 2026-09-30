package service

import (
	"fmt"
	"errors"
	"github.com/vanchaooo/go-marketplace/repository"
)

func Order(userID int, productID int) bool {
	if repository.UserExists(userID) {
		for _, v := range repository.Cart[userID] {
			if v.ProductID == productID {
				sum := repository.ProductCost(userID, productID)
				balance, _ := repository.GetBalance(userID)
				if balance < sum {
					repository.LowerBalance(userID, sum)
					return true
				} else {
					fmt.Println("На балансе не достаточно денег.")
					return false
				}
			}
		}
	}

	fmt.Println(errors.New("Не удалось оформить заказ."))
	return false
}

func CancelOrder(userID int, productID int) bool {
	if repository.UserExists(userID) {
		for _, v := range repository.GetOrders(userID) {
			if v.ProductID == productID {
				repository.DeleteOrder(userID, productID)
				repository.HigherBalance(userID, v.Price)
			}
		}
	}

	fmt.Println(errors.New("Не удалось отменить заказ."))
	return false
}

func GetMyOrders(userID int) {
	if repository.UserExists(userID) {
		repository.GetOrders(userID)
	}
}

func CheckOrdersHistory(userID int) bool {
	if repository.UserExists(userID) {
		repository.OrdersHistory(userID)
		return true
	}

	fmt.Println("Не удалось загрузить историю заказов.")
	return false
}