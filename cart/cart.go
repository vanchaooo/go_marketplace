package cart

import (
	"fmt"
	"errors"
	"github.com/vanchaooo/go-marketplace/catalog"
)

var Cart = make(map[string]int)

func AddToCart(category string, position string) bool {
	if category == "" || position == "" {
		fmt.Println(errors.New("Name can't be empty."))
		return false
	}

	if _, ok := catalog.Catalog[category][position]; ok {
		fmt.Printf("%s успешно добавлен в корзину!\n", position)
		Cart[position] = catalog.Catalog[category][position]
		return true
	} else {
		fmt.Println(errors.New("Incorrect name of category or position on catalog."))
		return false
	}
}

func GetCart() {
	fmt.Println("Товары в вашей корзине:")
	fmt.Println()
	for k, v := range Cart {
		fmt.Printf("Товар: %s, цена: %d\n", k, v)
	}
}

func DeleteFromCart(position string) bool {
	if position == "" {
		fmt.Println(errors.New("Name can't be empty."))
		return false
	}

	delete(Cart, position)
	return true
}