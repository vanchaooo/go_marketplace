package repository

import (
	"github.com/vanchaooo/go-marketplace/internal/model"
)

type WalletRepository struct {
	wallets map[int64]*model.Wallet
}

func NewWalletRepository() *WalletRepository {
	return &WalletRepository{
		wallets: make(map[int64]*model.Wallet),
	}
}

func (w WalletRepository) Add(wallet *model.Wallet) {
	w.wallets[wallet.UserID] = wallet
}

func (w WalletRepository) Get(userID int64) (*model.Wallet, bool) {
	wallet, ok := w.wallets[userID]
	return wallet, ok
}

func (w WalletRepository) Delete(userID int64) {
	delete(w.wallets, userID)
}
