package service

import (
	"context"
	"errors"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/vanchaooo/go-marketplace/internal/auth"
	"github.com/vanchaooo/go-marketplace/internal/model"
	"github.com/vanchaooo/go-marketplace/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type TokenPair struct {
	Access string
	Refresh string
}

type AuthService struct {
	users *repository.UserRepository
	wallets *repository.WalletRepository
	refresh *repository.RefreshTokenRepository
	tokens *auth.TokenManager
	refreshTTL time.Duration
}

func NewAuthService(
	users *repository.UserRepository,
	wallets *repository.WalletRepository,
	refresh *repository.RefreshTokenRepository,
	tokens *auth.TokenManager,
	refreshTTL time.Duration,
	) *AuthService {
		return &AuthService{
			users: users,
			wallets: wallets,
			refresh: refresh,
			tokens: tokens,
			refreshTTL: refreshTTL,
		}
}

func (a *AuthService) Register(ctx context.Context, name, email, password string) (model.User, error) {
	name = strings.TrimSpace(name)
	email = strings.ToLower(strings.TrimSpace(email))

	if length := utf8.RuneCountInString(name); length < 1 || length > 200 {
		return model.User{}, ErrInvalidName
	}
	if length := utf8.RuneCountInString(email); length < 1 || length > 200 {
		return model.User{}, ErrInvalidEmail
	}
	if len(password) < 8 || len(password) > 72 {
		return model.User{}, ErrInvalidPassword
	}

	passhash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}
	user, err := a.users.Add(ctx, name, email, string(passhash))
	if err != nil {
		if errors.Is(err, repository.ErrEmailTaken) {
			return model.User{}, ErrEmailTaken
		}
		return model.User{}, err
	}
	if _, err := a.wallets.Add(ctx, user.ID); err != nil {
		return model.User{}, err
	}
	
	return user, nil
}

func (a *AuthService) Login(ctx context.Context, email, password string) (TokenPair, error) {
	email = strings.ToLower(strings.TrimSpace(email))

	userID, hash, err := a.users.GetWithCredentials(ctx, email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return TokenPair{}, ErrInvalidCredentials
		}
		return TokenPair{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return TokenPair{}, ErrInvalidCredentials
	}
	return a.newTokenPair(ctx, userID)
}

func (a *AuthService) Logout(ctx context.Context, refreshToken string) error {
	return a.refresh.Delete(ctx, auth.HashToken(refreshToken))
}

func (a *AuthService) Refresh(ctx context.Context, refreshToken string) (TokenPair, error) {
	userID, err := a.refresh.Verification(ctx, auth.HashToken(refreshToken))
	if err != nil {
		if errors.Is(err, repository.ErrRefreshTokenNotFound) {
			return TokenPair{}, ErrInvalidToken
		}
		return TokenPair{}, err
	}
	return a.newTokenPair(ctx, userID)
}

func (a *AuthService) newTokenPair(ctx context.Context, userID int64) (TokenPair, error) {
	access, err := a.tokens.NewAccessToken(userID)
	if err != nil {
		return TokenPair{}, err
	}

	refresh, hash, err := auth.NewRefreshToken()
	if err != nil {
		return TokenPair{}, err
	}
	if err := a.refresh.Add(ctx, userID, hash, time.Now().Add(a.refreshTTL)); err != nil {
		return TokenPair{}, err
	}

	return TokenPair{Access: access, Refresh: refresh}, nil
}