package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager struct {
	Secret string
	TTL time.Duration
}

func NewTokenManaget(secret string, ttl time.Duration) *TokenManager {
	return &TokenManager{Secret: secret, TTL: ttl}
}

func (m *TokenManager) NewAccessToken(userID int64) (string, error) {
	now := time.Now()

	claims := jwt.RegisteredClaims {
		Subject: strconv.FormatInt(userID, 10),
		IssuedAt: jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(m.TTL)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.Secret)
}

func (m *TokenManager) ParseAccessToken(tokenString string) (int64, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {return m.Secret, nil}, jwt.WithValidMethods([]string{"HS256"}))
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, errors.New("токен недействителен")
	}

	return strconv.ParseInt(claims.Subject, 10, 64)
}

func HashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

func NewRefreshToken() (token string, hash string, err error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", "", err
	}

	token = base64.RawURLEncoding.EncodeToString(bytes)
	return token, HashToken(token), nil
}