package repository

import (
	"fmt"
	"errors"
	"slices"
	"time"
)

type Position struct {
	Name string
	Price int
	Amount int
}

type Data struct {
	Name string
	Time time.Time
}

var Orders = map[int][]*Position {}
var History = map[int][]*Data {}

func AddToOrders(userID int, product *Position) bool {
	if userID == 0 {
		fmt.Println(errors.New("Введите ваш ID."))
		return false
	}

	data := &Data {
		Name: product.Name,
		Time: time.Now(),
	}
	Orders[userID] = append(Orders[userID], product)
	History[userID] = append(History[userID], data)
	return true
}

func GetOrders(userID int) {
	for k, v := range Orders[userID] {
		fmt.Printf("%d. Товар: %s, Цена: %d, Кол-во: %d.", k, v.Name, v.Price, v.Amount)
	}
}

func DeleteOrder(userID int, position string) bool {
	if userID == 0 {
		fmt.Println(errors.New("Введите ваш ID."))
		return false
	}
	if position == "" {
		fmt.Println(errors.New("Имя товара не может быть пустым."))
		return false
	}

	for k, v := range Orders[userID] {
		if v.Name == position {
			Orders[userID] = slices.Delete(Orders[userID], k, k+1)
			fmt.Println("Заказ успешно отменен")
			return true
		}
	}

	fmt.Println("Не удалось отменить заказ.")
	return false
}

func OrdersHistory(userID int) {
	fmt.Println("Список всех ваших заказов: ")
	for _, v := range History {
		for _, j := range v {
			formatedTime := j.Time.Format("02.01.2006 в 15:04:05")
			fmt.Printf("Товар: %s, Дата заказа: %s.\n", j.Name, formatedTime)
		}
	}
}