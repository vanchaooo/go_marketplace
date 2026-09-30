package repository

import (
	"fmt"
	"errors"
	"slices"
	"time"
)

type Position struct {
	ProductID int
	Price int
	Amount int
}

type Data struct {
	ProductID int
	Time time.Time
}

var orders = map[int][]*Position {}
var history = map[int][]*Data {}

func AddToOrders(userID int, product *Position) bool {
	if UserExists(userID) {
		data := &Data {
			ProductID: product.ProductID,
			Time: time.Now(),
		}

		orders[userID] = append(orders[userID], product)
		history[userID] = append(history[userID], data)
		return true
	}

	fmt.Println(errors.New("Не удалось оформить заказ."))
	return false
}

func GetOrders(userID int) []*Position {
	if UserExists(userID) {
		result := []*Position{}
		for _, v := range orders[userID] {
			result = append(result, v)
			fmt.Printf("ID товара: %d, Цена: %d, Кол-во: %d.", v.ProductID, v.Price, v.Amount)
		}
		return result
	}
	return nil
}

func GetPrice(userID int, productID int) (int, bool) {
	if UserExists(userID) {
		if productID == 0 {
			fmt.Println(errors.New("ID продукта не может быть равно 0"))
			return 0, false
		}
		result := ProductCost(userID, productID)
		return result, true
	}

	fmt.Println(errors.New("Не удалось рассчитать сумму товаров."))
	return 0, false
}

func DeleteOrder(userID int, productID int) bool {
	prod, _ := GetProduct(productID)
	if UserExists(userID) {
		for k, v := range orders[productID] {
			if v.ProductID == prod.ID {
				orders[userID] = slices.Delete(orders[userID], k, k+1)
				fmt.Println("Заказ успешно отменен")
				return true
			}
		}
	}

	fmt.Println("Не удалось отменить заказ.")
	return false
}

func OrdersHistory(userID int) bool {
	if UserExists(userID) {
		fmt.Println("Список всех ваших заказов: ")
		for _, v := range history {
			for _, j := range v {
				formatedTime := j.Time.Format("02.01.2006 в 15:04:05")
				fmt.Printf("ID товара: %d, Дата заказа: %s.\n", j.ProductID, formatedTime)
				return true
			}
		}	
	}

	fmt.Println(errors.New("Не удалось загрузить историю заказов."))
	return false
}