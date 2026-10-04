package service

import (
	"fmt"
	"errors"
	"github.com/vanchaooo/go-marketplace/repository"
)

func Order(userID int64, productID int64) bool {
	if repository.UserExists(userID) {
		for _, v := range repository.Cart[userID] {
			if v.ProductID == productID {
				sum := repository.ProductCost(userID, productID)
				balance, _ := repository.GetBalance(userID)
				if balance > sum {
					item := &repository.Position{
						ProductID: productID,
						Amount: v.Amount,
					}
					repository.AddToOrders(userID, item)
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

func CancelOrder(userID int64, productID int64) bool {
	if repository.UserExists(userID) {
		for _, v := range repository.GetOrders(userID) {
			prod, _ := repository.GetProduct(productID)
			if v.ProductID == productID {
				repository.DeleteOrder(userID, productID)
				repository.HigherBalance(userID, prod.Price * v.Amount)
			}
		}
	}

	fmt.Println(errors.New("Не удалось отменить заказ."))
	return false
}

func GetMyOrders(userID int64) {
	if repository.UserExists(userID) {
		repository.GetOrders(userID)
	}
}

func CheckOrdersHistory(userID int64) bool {
	if repository.UserExists(userID) {
		repository.OrdersHistory(userID)
		return true
	}

	fmt.Println("Не удалось загрузить историю заказов.")
	return false
}