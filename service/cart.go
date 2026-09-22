package service

import (
	"github.com/vanchaooo/go-marketplace/repository"
)

func CheckCart(userID int) {
	if repository.UserExists(userID) {
		repository.GetCart(userID)
	}
}

func AddToCart(userID int, product string) {
	if repository.UserExists(userID) {
		repository.AddToCart(userID, product)
	}
}

func DeleteFromCart(userID int, product string) {
	if repository.UserExists(userID) {
		repository.DeleteFromCart(userID, product)
	}
}

func AddAmount(userID int, product string) {
	if repository.UserExists(userID) {
		repository.AddAmount(userID, product)
	}
}

func LowAmount(userID int, product string) {
	if repository.UserExists(userID) {
		repository.LowerAmount(userID, product)
	}
}