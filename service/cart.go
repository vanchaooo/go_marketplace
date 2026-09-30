package service

import (
	"github.com/vanchaooo/go-marketplace/repository"
)

func CheckCart(userID int) {
	if repository.UserExists(userID) {
		repository.GetCart(userID)
	}
}

func AddToCart(userID int, productID int) {
	if repository.UserExists(userID) {
		repository.AddToCart(userID, productID)
	}
}

func DeleteFromCart(userID int, productID int) {
	if repository.UserExists(userID) {
		repository.DeleteFromCart(userID, productID)
	}
}

func AddAmount(userID int, productID int) {
	if repository.UserExists(userID) {
		repository.AddAmount(userID, productID)
	}
}

func LowAmount(userID int, productID int) {
	if repository.UserExists(userID) {
		repository.LowerAmount(userID, productID)
	}
}