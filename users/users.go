package users

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"github.com/vanchaooo/go-marketplace/balance"
)

var (
	Users_list = map[int]string {
		1: "Jack",
		2: "Alice",
		3: "Johny",
		4: "Bob",
		5: "Katty",
		6: "Kevin",
	}

	count = len(Users_list)+1
)

func UserExists(userID int) bool {
	if _, ok := Users_list[userID]; !ok {
		fmt.Println(errors.New("User does not exsist."))
		return false
	} else {
		return true
	}
}

func AddUser(name string) int {
	if name == "" {
		fmt.Println(errors.New("Name can't be empty."))
		return 0
	}

	Users_list[count] = name
	return count
}

func GetUser(id int) (string, bool) {
	if id <= 0 {
		fmt.Println(errors.New("ID should be more than 0."))
		return "", false
	}

	if v, ok := Users_list[id]; ok {
		return v, true
	} else {
		fmt.Println(errors.New("Unknown user."))
		return v, false
	}
}

func GetAllusers() {
	for k, v := range Users_list {
		fmt.Printf("ID: %d, Пользователь: %s\n", k, v)
	}
}

func RenameUser(id int, newName string) bool {
	if id <= 0 || id > len(Users_list) {
		fmt.Println(errors.New("Incorrect ID."))
		return false
	}

	if newName == "" {
		fmt.Println(errors.New("New name can't be empty."))
		return false
	}


	if _, ok := Users_list[id]; ok {
		Users_list[id] = newName
		return true
	} else {
		fmt.Println("Invalid user.")
		return false
	}
}

func DeleteUser(id int) bool {
	if id <= 0 || id > len(Users_list) {
		fmt.Println(errors.New("Incorrect ID."))
		return false
	}

	delete(Users_list, id)
	delete(balance.Balances, id)
	return true
}

func FindUsersByName(query string) []int {
	if query == "" {
		fmt.Println(errors.New("Name can't be empty."))
		return nil
	}

	result := []int{}
	for k, v := range Users_list {
		if strings.EqualFold(v, query) {
			result = append(result, k)
		}
	}

	slices.Sort(result)
	return result
}