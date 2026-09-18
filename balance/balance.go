package balance

import (
	"fmt"
	"errors"
)

var (
	Balances = map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
	}
)

func SetBalance(
userID int,
amount int,
) bool {
	if amount <= 0 {
		fmt.Println(errors.New("Amount can't be negative or zero."))
		return false
	}

	Balances[userID] += amount
	return true
}

func GetEveryoneBalance() {
	for k, v := range Balances {
		rubles := v / 100
		kopecks := v % 100
		fmt.Printf("ID: %d, Баланс: %d рублей %d копеек.\n", k, rubles, kopecks)
	}
}

func GetBalance(
users map[int]string,
userID int,
) (int, bool) {
	if _, ok := users[userID]; !ok {
		fmt.Println(errors.New("User does not exsist."))
		return 0, false
	}

	rubles := Balances[userID] / 100
	kopecks := Balances[userID] % 100

	fmt.Printf("Пользователь (ID): %d. Баланс: %d рублей %d копеек\n", userID, rubles, kopecks)
	result, ok := Balances[userID]
	return result, ok
}