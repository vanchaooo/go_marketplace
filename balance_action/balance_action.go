package balance_action

import (
	"fmt"
	"errors"
)

func TopUpBalance(
users map[int]string,
balances map[int]int,
userID int,
amount int,
) bool {
	if _, ok := users[userID]; !ok {
		fmt.Println(errors.New("User does not exsist."))
		return false
	}

	if amount <= 0 {
		fmt.Println(errors.New("Amount can't be negative or zero."))
		return false
	}

	balances[userID] += amount
	return true
}

func GetBalance(
users map[int]string,
balances map[int]int,
userID int,
) (int, bool) {
	if _, ok := users[userID]; !ok {
		fmt.Println(errors.New("User does not exsist."))
		return 0, false
	}

	rubles := balances[userID] / 100
	kopecks := balances[userID] % 100

	fmt.Printf("Пользователь (ID): %d. Баланс: %d рублей %d копеек\n", userID, rubles, kopecks)
	result, ok := balances[userID]
	return result, ok
}