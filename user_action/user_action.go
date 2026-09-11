package user_action

import (
	"errors"
	"fmt"
	"strings"
	"slices"
	// "github.com/vanchaooo/go-marketplace/balance_action"
)

func AddUser(users map[int]string, balances map[int]int, id int, name string) bool {
	if id <= 0 {
		fmt.Println(errors.New("ID should be more than 0."))
		return false
	}

	if name == "" {
		fmt.Println(errors.New("Name can't be empty."))
		return false
	}

	if _, ok := users[id]; ok {
		fmt.Println(errors.New("User alredy exsist."))
		return false
	} else {
		users[id] = name
		balances[id] = 0
		return true
	}
}

func GetUser(users map[int]string, id int) (string, bool) {
	if id <= 0 {
		fmt.Println(errors.New("ID should be more than 0."))
		return "", false
	}

	if v, ok := users[id]; ok {
		return v, true
	} else {
		fmt.Println(errors.New("Unknown user."))
		return v, false
	}
}

func RenameUser(users map[int]string, id int, newName string) bool {
	if id <= 0 || id > len(users) {
		fmt.Println(errors.New("Incorrect ID."))
		return false
	}

	if newName == "" {
		fmt.Println(errors.New("New name can't be empty."))
		return false
	}


	if _, ok := users[id]; ok {
		users[id] = newName
		return true
	} else {
		fmt.Println("Invalid user.")
		return false
	}
}

func DeleteUser(users map[int]string, balances map[int]int, id int) bool {
	if id <= 0 || id > len(users) {
		fmt.Println(errors.New("Incorrect ID."))
		return false
	}

	delete(users, id)
	delete(balances, id)
	return true
}

func FindUsersByName(users map[int]string, query string) []int {
	if query == "" {
		fmt.Println(errors.New("Name can't be empty."))
		return nil
	}

	result := []int{}

	for k, v := range users {
		if strings.EqualFold(v, query) {
			result = append(result, k)
		}
	}

	slices.Sort(result)

	return result
}