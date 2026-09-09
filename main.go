package main

import (
	"fmt"
	"github.com/vanchaooo/go-marketplace/user_action"
	"github.com/vanchaooo/go-marketplace/balance_action"
)



func main () {
	users := map[int]string{}
	balances := map[int]int{}

	
	user_action.AddUser(users, balances, 1, "Ivan")
	user_action.AddUser(users, balances, 2, "Bob")
	user_action.AddUser(users, balances, 3, "John")
	user_action.AddUser(users, balances, 4, "Alice")
	user_action.AddUser(users, balances, 5, "Zoi")
	user_action.AddUser(users, balances, 6, "Bob")

	for k, v := range users {
		fmt.Printf("Ключ: %d, Значение: %s\n", k, v)
	}


	balance_action.TopUpBalance(users, balances, 4, 1000)
	balance_action.GetBalance(users, balances, 4)

	// for k, v := range balances {
	// 	fmt.Printf("Ключ: %d, Значение: %d\n", k, v)
	// }
}