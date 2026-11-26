package service

import (
	"context"
	"errors"
	"github.com/vanchaooo/go-marketplace/internal/model"
	"github.com/vanchaooo/go-marketplace/internal/repository"
)

type WalletService struct {
	users *repository.UserRepository
	wallets *repository.WalletRepository
}

func NewWalletService(wallets *repository.WalletRepository, users *repository.UserRepository) *WalletService {
	return &WalletService{
		users: users,
		wallets: wallets,
	}
}

func (w *WalletService) TopUpBalance(ctx context.Context, userID, amount int64) (model.Wallet, error) {
	exist, err := w.users.Exists(ctx, userID)
	if !exist {
		return model.Wallet{}, ErrUserNotFound
	}
	if err != nil {
		return model.Wallet{}, err
	}
	if amount <= 0 {
		return model.Wallet{}, ErrAmountNegative
	}

	updWallet, err := w.wallets.TopUp(ctx, userID, amount)
	if err != nil {
		return model.Wallet{}, err
	}
	return updWallet, nil
}

func (w *WalletService) WithdrawBalance(ctx context.Context, userID, amount int64) (model.Wallet, error) {
	if amount <= 0 {
		return model.Wallet{}, ErrAmountNegative
	}

	exist, err := w.users.Exists(ctx, userID)
	if err != nil {
		return model.Wallet{}, err
	}
	if !exist {
		return model.Wallet{}, ErrUserNotFound
	}

	updWallet, err := w.wallets.Withdraw(ctx, userID, amount)
	if err != nil {
		if errors.Is(err, repository.ErrInsufficientFunds) {
			return model.Wallet{}, ErrInsufficientFunds
		}
		if errors.Is(err, repository.ErrWalletNotFound) {
			return model.Wallet{}, ErrWalletNotFound
		}
		return model.Wallet{}, err
	}
	return updWallet, nil
}

func (w *WalletService) GetWallet(ctx context.Context, userID int64) (model.Wallet, error) {
	exist, err := w.users.Exists(ctx, userID)
	if err != nil {
		return model.Wallet{}, err
	}
	if !exist {
		return model.Wallet{}, ErrUserNotFound
	}

	wallet, err := w.wallets.Get(ctx, userID)
	if err != nil {
		if errors.Is(err, repository.ErrWalletNotFound) {
			return model.Wallet{}, ErrWalletNotFound
		}
		return model.Wallet{}, err
	}
	return wallet, nil
}