package repository

import (
	"errors"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vanchaooo/go-marketplace/internal/model"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository {
		pool: pool,
	}
}

func (u *UserRepository) Add(ctx context.Context, name string) (model.User, error) {
	const query = `
		INSERT INTO users (name)
		VALUES ($1)
		RETURNING id, name, created_at
	`

	var user model.User
	if err := u.pool.QueryRow(ctx, query, name).Scan(
		&user.ID,
		&user.Name,
		&user.CreatedAt,
	); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *UserRepository) Get(ctx context.Context, userID int64) (model.User, error) {
	const query = `
		SELECT id, name, created_at
		FROM users
		WHERE id = $1
	`
	var user model.User
	if err := u.pool.QueryRow(ctx, query, userID).Scan(
		&user.ID,
		&user.Name,
		&user.CreatedAt,
	); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *UserRepository) Rename(ctx context.Context, userID int64, newName string) (model.User, error) {
	const query = `
		UPDATE users
		SET name = $2
		WHERE id = $1
		RETURNING id, name, created_at
	`
	var renamedUser model.User
	if err := u.pool.QueryRow(ctx, query, userID, newName).Scan(
		&renamedUser.ID,
		&renamedUser.Name,
		&renamedUser.CreatedAt,
	); err != nil {
		return model.User{}, err
	}
	return renamedUser, nil
}

func (u *UserRepository) Delete(ctx context.Context, userID int64) error {
	const query = `
		DELETE FROM users
		WHERE id = $1
	`
	response, err := u.pool.Exec(ctx, query, userID)
	if err != nil {
		return err
	}
	if response.RowsAffected() == 0 {
		return errors.New("Пользоватесь с таким ID не найден.")
	}
	return nil
}

func (u *UserRepository) Exists(ctx context.Context, userID int64) (bool, error) {
	const query = `
		SELECT EXISTS (
			SELECT id
			FROM users
			WHERE id = $1
		)
	`
	var exist bool
	if err := u.pool.QueryRow(ctx, query, userID).Scan(&exist); err != nil {
		return false, err
	}
	return exist, nil
}