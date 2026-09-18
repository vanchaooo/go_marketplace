package repository

import (
	"fmt"
	"errors"
)

var Balances = map[int]int{
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

func CreateBalance(userID int) {
	Balances[userID] = 0
}

func GetBalance(userID int) (int, bool) {
	if balance, ok := Balances[userID]; ok {
		if userID <= 0 {
			fmt.Println(errors.New("ID Пользователя не может быть меньше 0."))
			return 0, false
		}

		fmt.Printf("ID пользователя: %d. Баланс: %d", userID, balance)
		return balance, ok
	} else {
		fmt.Println(errors.New("Не удалось найти пользователя с таким ID."))
		return 0, false
	}
}

func SetBalance(userID int, balance int) bool {
	if _, ok := Balances[userID]; ok {
		if userID <= 0 {
			fmt.Println(errors.New("ID Пользователя не может быть меньше 0."))
			return false
		}
		if balance <= 0 {
			fmt.Println(errors.New("Устанавливаемый баланс не может быть меньше 0."))
			return false
		}

		Balances[userID] = balance
		return true
	}

	fmt.Println(errors.New("ПОльзователь с таким ID не найден."))
	return false
}

func DeleteBalance(userID int) bool {
	if _, ok := Balances[userID]; ok {
		if userID <= 0 {
			fmt.Println(errors.New("ID Пользователя не может быть меньше 0."))
			return false
		}

		delete(Balances, userID)
		return true
	}

	fmt.Println(errors.New("Пользователь с таким ID не найден."))
	return false
}