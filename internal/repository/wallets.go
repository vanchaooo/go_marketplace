package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vanchaooo/go-marketplace/internal/model"
)

var ErrWalletNotFound = errors.New("кошелек не найден")
var ErrInsufficientFunds = errors.New("недостаточно средств")

type WalletRepository struct {
	pool *pgxpool.Pool
}

func NewWalletRepository(pool *pgxpool.Pool) *WalletRepository {
	return &WalletRepository{pool: pool}
}

func (b *WalletRepository) Add(ctx context.Context, userID int64) (model.Wallet, error) {
	const query = `
		INSERT INTO wallets (user_id)
		VALUES ($1)
		RETURNING user_id, balance
	`

	var wallet model.Wallet
	if err := b.pool.QueryRow(ctx, query, userID).Scan(
		&wallet.UserID,
		&wallet.Balance,
	); err != nil {
		return model.Wallet{}, err
	}
	return wallet, nil
}

func (b *WalletRepository) TopUpBalance(ctx context.Context, userID, amount int64) (model.Wallet, error) {
	const query = `
		UPDATE wallets 
		SET balance = balance + $2
		WHERE user_id = $1
		RETURNING user_id, balance
	`
	var wallet model.Wallet
	if err := b.pool.QueryRow(ctx, query, userID, amount).Scan(
		&wallet.UserID,
		&wallet.Balance,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Wallet{}, ErrWalletNotFound
		}
		return model.Wallet{}, err
	}
	return wallet, nil
}

func (b *WalletRepository) Withdraw(ctx context.Context, userID, amount int64) (model.Wallet, error) {
	const query = `
		UPDATE wallets 
		SET balance = balance - $2
		WHERE user_id = $1 AND balance >= $2
		RETURNING user_id, balance
	`
	var wallet model.Wallet
	err := b.pool.QueryRow(ctx, query, userID, amount).Scan(&wallet.UserID, &wallet.Balance)
	if err == nil {
		return wallet, nil
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return model.Wallet{}, err
	}

	var exists bool
	if err := b.pool.QueryRow(ctx, `SELECT EXISTS (SELECT 1 FROM wallets WHERE user_id = $1)`, userID).Scan(&exists);  
	err != nil {
		return model.Wallet{}, err
	}
	if !exists {
		return model.Wallet{}, ErrWalletNotFound
	}
	return model.Wallet{}, ErrInsufficientFunds
}

func (b *WalletRepository) Get(ctx context.Context, userID int64) (model.Wallet, error) {
	const query = `
		SELECT user_id, balance
		FROM wallets
		WHERE user_id = $1
	`
	var wallet model.Wallet
	if err := b.pool.QueryRow(ctx, query, userID).Scan(
		&wallet.UserID,
		&wallet.Balance,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Wallet{}, ErrWalletNotFound
		}
		return model.Wallet{}, err
	}
	return wallet, nil
}

func (b *WalletRepository) Delete(ctx context.Context, userID int64) error {
	const query = `
		DELETE FROM wallets
		WHERE user_id = $1
	`
	response, err := b.pool.Exec(ctx, query, userID)
	if err != nil {
		return err
	}
	if response.RowsAffected() == 0 {
		return ErrWalletNotFound
	}
	return nil
}