package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vanchaooo/go-marketplace/internal/model"
)

var (
	ErrEmailTaken = errors.New("email уже занят")
	ErrUserNotFound = errors.New("Пользователь не найден.")
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository {
		pool: pool,
	}
}

func (u *UserRepository) Add(ctx context.Context, name, email, password_hash string) (model.User, error) {
	const query = `
		INSERT INTO users (name, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, name, created_at
	`

	var user model.User
	if err := u.pool.QueryRow(ctx, query, name, email, password_hash).Scan(
		&user.ID,
		&user.Name,
		&user.CreatedAt,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return model.User{}, ErrEmailTaken
		}
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

func (u *UserRepository) GetWithCredentials(ctx context.Context, email string) (int64, string, error) {
	const query = `
		SELECT id
		FROM users
		WHERE email = $1
	`
	var (
		id int64
		hash string
	)
	if err := u.pool.QueryRow(ctx, query, email).Scan(&id, &hash); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, "", ErrUserNotFound
		}
		return 0, "", err
	}
	return id, hash, nil
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