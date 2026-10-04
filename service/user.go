package service

import (
	"fmt"
	"github.com/vanchaooo/go-marketplace/repository"
)

func RegisterUser(name string) {
	strct := &repository.User{
		Name: name,
	}
	fmt.Println("Регистрация прошла успешно!")
	repository.AddUser(strct)
	SetZeroBalance(strct.ID)
}

func SetZeroBalance(userID int64) {
	if repository.UserExists(userID) {
		repository.CreateBalance(userID)
	}
}

func UpdateBalance(userID int64, number int64) {
	if repository.UserExists(userID) {
		repository.HigherBalance(userID, number)
	}
}

func LowerBalance(userID int64, number int64) {
	if repository.UserExists(userID) {
		repository.LowerBalance(userID, number)
	}
}