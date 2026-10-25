package service

import (
	"github.com/vanchaooo/go-marketplace/internal/model"
	"github.com/vanchaooo/go-marketplace/internal/repository"
)

type UserService struct {
	users *repository.UserRepository
	wallets *repository.WalletRepository
	carts *repository.CartRepository
	orders *repository.OrderRepository
}

func NewUserService(
	users *repository.UserRepository,
	wallets *repository.WalletRepository,
	carts *repository.CartRepository,
	orders *repository.OrderRepository,	
) *UserService {
	return &UserService{
		users: users,
		wallets: wallets,
		carts: carts,
		orders: orders,
	}
}

func (u *UserService) AddUser(name string) (int64, bool) {
	user, ok := model.NewUser(name)
	if !ok {
		return 0, false
	}

	userID, ok := u.users.Add(&user)
	if !ok {
		return 0, false
	}

	wallet, ok := model.NewWallet(userID)
	if !ok {
		return 0, false
	}

	cart, ok := model.NewCart(userID)
	if !ok {
		return 0, false
	}

	u.wallets.Add(&wallet)
	u.carts.Add(&cart)
	return userID, true
}

func (u *UserService) Renameuser(userID int64, newName string) bool {
	user, ok := u.users.Get(userID)
	if !ok {
		return false
	}

	return user.Rename(newName)
}

func (u *UserService) GetUser(userID int64) (model.User, bool) {
	user, ok := u.users.Get(userID)
	if !ok {
		return model.User{}, false
	}

	return *user, true
}

func (u *UserService) DeleteUser(userID int64) bool {
	if !u.users.Exists(userID) {
		return false
	}

	wallet, ok := u.wallets.Get(userID)
	if !ok || wallet.Balance != 0 {
		return false
	}
 
	for _, order := range u.orders.GetAll(userID) {
		if order.IsPaid() {
			return false
		}
	}
 
	u.users.Delete(userID)
	u.wallets.Delete(userID)
	u.carts.Delete(userID)
	return true
}

func (u *UserService) FindEveryoneByName(name string) []int64 {
	return u.users.GetEveryoneByName(name)
}