package orders

import (
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/vanchaooo/go-marketplace/cart"
)

var (
	AllOrders = []string{}

	History = make(map[string]time.Time)
)

func Order(position string) bool {
	if position == "" {
		fmt.Println(errors.New("Неверное наименование товара."))
		return false
	}

	if _, ok := cart.Cart[position]; ok {
		fmt.Printf("Заказ успешно оформлен! Товар %s был удален из корзины и перемещен в заказы.\n", position)
		AllOrders = append(AllOrders, position)
		History[position] = time.Now()
		delete(cart.Cart, position)
		return true
	} else {
		fmt.Printf("В вашей корзине нету %s.\n", position)
		return false
	}
}

func GetAllOrders() {
	fmt.Println("Все ваши заказы:")
	fmt.Println()
	for _, v := range AllOrders {
		fmt.Println(v)
	}
}

func CancelOrder(position string) bool {
	if position == "" {
		fmt.Println(errors.New("Неверное наименование товара."))
		return false
	}

	exists := slices.Contains(AllOrders, position)
	indx := slices.Index(AllOrders, position)
	if exists {
		AllOrders = slices.Delete(AllOrders, indx, indx+1)
		fmt.Printf("Заказ товара %s успешно отменен!\n", position)
		return true
	} else {
		fmt.Println(errors.New("Товар не найден!"))
		return false
	}
}

func OrdersHistory() {
	fmt.Println("Список всех ваших заказов: ")
	for k, v := range History {
		formatedTime := v.Format("02.01.2006 в 15:04:05")
		fmt.Printf("Предмет: %s, Дата заказа: %s.\n", k, formatedTime)
	}
}