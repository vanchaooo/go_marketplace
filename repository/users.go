package repository

import (
	"fmt"
	"errors"
)

type User struct {
	ID   int
	Name string
}

var Users = map[int]*User{
	1:  {ID: 1, Name: "Иван"},
	2:  {ID: 2, Name: "Мария"},
	3:  {ID: 3, Name: "Алексей"},
	4:  {ID: 4, Name: "Ольга"},
	5:  {ID: 5, Name: "Дмитрий"},
	6:  {ID: 6, Name: "Елена"},
	7:  {ID: 7, Name: "Александр"},
	8:  {ID: 8, Name: "Анна"},
	9:  {ID: 9, Name: "Михаил"},
	10: {ID: 10, Name: "Татьяна"},
	11: {ID: 11, Name: "Сергей"},
	12: {ID: 12, Name: "Наталья"},
	13: {ID: 13, Name: "Артем"},
	14: {ID: 14, Name: "Ирина"},
	15: {ID: 15, Name: "Николай"},
}

func UserExists(userID int) bool {
	if _, ok := Users[userID]; !ok {
		fmt.Println(errors.New("Не удалось найти пользователя с таким ID."))
		return false
	} else {
		return true
	}
}

func AddUser(user *User) {
	if user.Name == "" {
		fmt.Println(errors.New("Имя пользователя не может быть пустым."))
	}

	user.ID = len(Users)+1
	Users[user.ID] = user
}

func GetAllUsers() {
	for _, v := range Users {
		fmt.Printf("ID: %d, Имя: %s.\n", v.ID, v.Name)
	}
}

func GetUser(id int) (*User, bool) {
	if id <= 0 {
		fmt.Println(errors.New("ID Пользователя должно быть больше 0."))
		return nil, false
	}
	
	if user, ok := Users[id]; ok {
		fmt.Printf("По такому ID был найден: %q.\n", user.Name)
		return user, ok
	}

	fmt.Println(errors.New("Пользователь с таким ID не был найден."))
	return nil, false
}

func DeleteUser(id int) bool {
	if id <= 0 {
		fmt.Println(errors.New("ID Пользователя должно быть больше 0."))
		return false
	}

	if _, ok := Users[id]; ok {
		fmt.Println("Пользователь успешно удален.")
		delete(Users, id)
		return true
	}
	
	fmt.Println(errors.New("Пользователя с таким ID не найдено."))
	return false
}