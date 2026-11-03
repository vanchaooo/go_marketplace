package service

import (
	"context"
	"unicode/utf8"
	"github.com/vanchaooo/go-marketplace/internal/model"
	"github.com/vanchaooo/go-marketplace/internal/repository"
)

type UserService struct {
	users *repository.UserRepository
}

func NewUserService(users *repository.UserRepository) *UserService {
	return &UserService{
		users: users,
	}
}

func (u *UserService) AddUser(ctx context.Context, name string) (model.User, error) {
	if length := utf8.RuneCountInString(name); length < 1 || length > 200 {
		return model.User{}, ErrInvalidName
	}
	response, err := u.users.Add(ctx, name)
	if err != nil {
		return model.User{}, err
	}
	return response, nil
}

func (u *UserService) RenameUser(ctx context.Context, userID int64, newName string) (model.User, error) {
	if userID <= 0 {
		return model.User{}, ErrInvalidID
	}
	if length := utf8.RuneCountInString(newName); length < 1 || length > 200 {
		return model.User{}, ErrInvalidName
	}
	exist, err := u.users.Exists(ctx, userID)
	if err != nil {
		return model.User{}, err
	}
	if !exist {
		return model.User{}, ErrUserNotFound
	}

	renamedUser, err := u.users.Rename(ctx, userID, newName)
	if err != nil {
		return model.User{}, nil
	}
	return renamedUser, nil
}

func (u *UserService) GetUser(ctx context.Context, userID int64) (model.User, error) {
	if userID <= 0 {
		return model.User{}, ErrInvalidID
	}
	exist, err := u.users.Exists(ctx, userID)
	if err != nil {
		return model.User{}, err
	}
	if !exist {
		return model.User{}, ErrUserNotFound
	}

	user, err := u.users.Get(ctx, userID)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *UserService) DeleteUser(ctx context.Context, userID int64) error {
	if userID <= 0 {
		return ErrInvalidID
	}
	exist, err := u.users.Exists(ctx, userID)
	if err != nil {
		return err
	}
	if !exist {
		return ErrUserNotFound
	}

	if err := u.users.Delete(ctx, userID); err != nil {
		return err
	}
	return nil
}