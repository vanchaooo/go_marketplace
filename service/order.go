package service

import (
	"fmt"
	"errors"
	"github.com/vanchaooo/go-marketplace/repository"
)

func Order(userID int, position string) bool {
	if repository.UserExists(userID) {
		for _, v := range repository.Cart[userID] {
			if v.Name == position {
				sum := v.Price
				balance, _ := repository.GetBalance(userID)
				if balance >= sum {
					tovar := &repository.Position{
						Name: v.Name,
						Price: v.Price,
						Amount: v.Amount,
					}
					fmt.Println("Заказ успешно офрмлен!")
					repository.AddToOrders(userID, tovar)
					repository.DeleteFromCart(userID, position)
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

func CancelOrder(userID int, position string) bool {
	if repository.UserExists(userID) {
		for _, v := range repository.Orders[userID] {
			if v.Name == position {
				repository.DeleteOrder(userID, position)
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