package repository

import (
	"errors"
	"fmt"
	"slices"
	"time"
)

type Position struct {
	ProductID int64
	Amount int64
}

type Data struct {
	ProductID int64
	Time time.Time
}

var orders = map[int64][]*Position {}
var history = map[int64][]*Data {}

func AddToOrders(userID int64, product *Position) bool {
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

func GetOrders(userID int64) []*Position {
	if UserExists(userID) {
		result := []*Position{}

		if len(orders) == 0 {
			fmt.Println("Вы еще ничего не заказали")
			return nil
		}

		for _, v := range orders[userID] {
			prod, _ := GetProduct(v.ProductID)
			result = append(result, v)
			fmt.Printf("ID товара: %d, Товар: %s, Цена: %d, Кол-во: %d.\n", v.ProductID, prod.Name, prod.Price * v.Amount, v.Amount)
		}
		return result
	}
	return nil
}

func GetPrice(userID int64, productID int64) (int64, bool) {
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

func DeleteOrder(userID int64, productID int64) bool {
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

func OrdersHistory(userID int64) bool {
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