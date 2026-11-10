package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrRefreshTokenNotFound = errors.New("refresh токен не найден")

type RefreshTokenRepository struct {
	pool *pgxpool.Pool
}

func NewRefreshTokenRepository(pool *pgxpool.Pool) *RefreshTokenRepository {
	return &RefreshTokenRepository{pool: pool}
}

func (t *RefreshTokenRepository) Add(ctx context.Context, userID int64, hash string, exp time.Time) error {
	const query = `
		INSERT INTO refresh_tokens (user_id, token_hash, expires_at)
		VALUES ($1, $2, $3)
	`

	_, err := t.pool.Exec(ctx, query, userID, hash, exp)
	return err
}

func (t *RefreshTokenRepository) Verification(ctx context.Context, userID int64, hash string) (int64, error) {
	const query = `
		DELETE FROM refresh_tokens
		WHERE token_hash = $1 and expires_at > now()
		RETURNING user_id
	`

	var ID int64
	if err := t.pool.QueryRow(ctx, query, hash).Scan(&ID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, ErrRefreshTokenNotFound
		}
		return 0, nil
	}
	return ID, nil
}

func (t *RefreshTokenRepository) Delete(ctx context.Context, hash string) error {
	const query = `
		DELETE FROM refresh_tokens
		WHERE token_hash = $1
	`

	_, err := t.pool.Exec(ctx, query, hash)
	return err
}