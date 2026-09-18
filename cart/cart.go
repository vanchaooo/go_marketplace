package cart

import (
	"errors"
	"fmt"
	"github.com/vanchaooo/go-marketplace/catalog"
)

var (
	Cart = map[string]map[int]int{}
)

// func get_quantity() int {
// 	count := 0 
// 	for _, v := range Cart {
// 		for _, amount := range v {
// 			count += amount
// 		}
// 	}
// 	return count
// }


// Добавление товара в корзину
func AddToCart(category string, position string, product_quantity int) bool {
	if category == "" || position == "" {
		fmt.Println(errors.New("Пожалуйста, проверьте, заполнили ли вы поля."))
		return false
	}

	product, ok := catalog.Catalog[category][position]

	if ok {
		for price, stock := range product {
			if product_quantity <= 0 {
				fmt.Println(errors.New("Количество товара должно быть больше 0."))
                return false
			}

			if product_quantity > stock {
				fmt.Println(errors.New("На складе недостаточно товара."))
                return false
			}

			Cart[position] = map[int]int {
				price: product_quantity,
			}

			fmt.Printf("%s успешно добавлен в корзину!\n", position)
			return true
		}
	}

	fmt.Printf("Не удалось добавить %q в корзину.\n", position)
	return false
}


// Получения всей корзины
func GetCart() {
	total := 0
    for _, product := range Cart {
        for price, amount := range product {
            total += price * amount
        }
    }

	fmt.Printf("В корзине %d предметов на %d рублей: ", get_quantity(), total)
	fmt.Println()
	for k, v := range Cart {
		for price, stock := range v {
			fmt.Printf("Товар: %s. Цена: %d, Количество в корзине: %d\n", k, price, stock)
		}
	}
}


// Полное удаление товара из корзины
func DeleteFromCart(position string) bool {
	if position == "" {
		fmt.Println(errors.New("Пожалуйста, проверьте, заполнили ли вы поля."))
		return false
	}

	delete(Cart, position)
	fmt.Printf("Успешное удаление %q из корзины товаров!\n", position)
	return true
}


// Уменьшение количества товара в корзине
func ReduceQuantity(position string) bool {
	if position == "" {
		fmt.Println(errors.New("Пожалуйста, проверьте, заполнили ли вы поля."))
		return false
	}

	product, ok := Cart[position]
	if ok {
		for price, stock := range product {
			if stock > 1 {
				Cart[position] = map[int]int {
					price: stock-1,
				}
			} else {
				DeleteFromCart(position)
				return true
			}
		}
		fmt.Printf("Количество %q успешно уменьшено на 1 шт.\n", position)
		return true
	}

	fmt.Println("Не удалось уменьшить количество товара.")
	return false
}


// Увеличение количества товара в корзине
func AddQuantity(position string) bool {
	if position == "" {
		fmt.Println(errors.New("Пожалуйста, проверьте, заполнили ли вы поля."))
		return false
	}

	product, ok := Cart[position]
	if ok {
			for price, stock := range product {
				Cart[position] = map[int]int {
					price: stock+1,
				}
			}
			fmt.Printf("Количество %q успешно увеличено на 1 шт.\n", position)
			return true
	}

	fmt.Println("Не удалось увеличить количество товара.")
	return false
}