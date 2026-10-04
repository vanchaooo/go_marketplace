package repository

import (
	"fmt"
	"errors"
)

var balances = map[int64]int64{
	1: 0,
	2: 0,
	3: 0,
	4: 0,
	5: 0,
	6: 0,
	7: 0,
	8: 0,
	9: 0,
	10: 0,
	11: 0,
	12: 0,
	13: 0,
	14: 0,
	15: 0,
}

func CreateBalance(userID int64) {
	balances[userID] = 0
}

func GetEveryoneBalance() {
	for k, v := range balances {
		fmt.Printf("ID: %d, Баланс: %d\n", k, v)
	}
}

func GetBalance(userID int64) (int64, bool) {
	if UserExists(userID) {
		fmt.Printf("ID пользователя: %d. Баланс: %d\n", userID, balances[userID])
		return balances[userID], true
	}

	fmt.Println(errors.New("Не удалось найти пользователя с таким ID."))
	return 0, false
}

func DeleteBalance(userID int64) bool {
	if UserExists(userID) {
		delete(balances, userID)
		fmt.Println("Баланс удален.")
		return true
	}

	fmt.Println(errors.New("Пользователь с таким ID не найден."))
	return false
}

func LowerBalance(userID int64, balance int64) bool {
	if UserExists(userID) {
		if balance == 0 {
			fmt.Println(errors.New("Сумма не может быть равна 0"))
			return false
		}
		if balance >= balances[userID] {
			DeleteBalance(userID)
		}

		balances[userID] -= balance
		fmt.Printf("Баланс уменьшен на %d рублей.\n", balance)
		fmt.Printf("Текущий баланс: %d\n", balances[userID])
		return true
	}

	fmt.Println(errors.New("Ошибка."))
	return false
}

func HigherBalance(userID int64, balance int64) bool {
	if UserExists(userID) {
		if balance == 0 {
			fmt.Println(errors.New("Сумма не может быть равна 0."))
			return false
		}

		balances[userID] += balance
		fmt.Printf("Баланс увеличен на %d рублей.\n", balance)
		fmt.Printf("Текущий баланс: %d\n", balances[userID])
		return true
	}

	fmt.Println(errors.New("Ошибка."))
	return false
}