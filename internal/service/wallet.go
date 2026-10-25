package service

import (
	"github.com/vanchaooo/go-marketplace/internal/repository"
)

type WalletService struct {
	users *repository.UserRepository
	wallets *repository.WalletRepository
}

func NewWalletService(
	users *repository.UserRepository,
	wallets *repository.WalletRepository,
) *WalletService {
	return &WalletService {
		users: users,
		wallets: wallets,
	}
}

func (w *WalletService) TopUpBalance(userID, amount int64) bool {
	if !w.users.Exists(userID) {
		return false
	}

	wallet, ok := w.wallets.Get(userID)
	if !ok {
		return false
	}

	return wallet.TopUp(amount)
}

func (w *WalletService) GetBalance(userID int64) (int64, bool) {
	wallet, ok := w.wallets.Get(userID)
	if !ok {
		return 0, false
	}

	return wallet.Balance, true
}