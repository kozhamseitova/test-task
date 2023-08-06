package jwttoken

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("invalid token err")
	ErrExpiredToken = errors.New("expired token err")
)

type JWTPayload struct {
	ID        uuid.UUID `json:"id"`
	UserId    int     `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (j JWTPayload) Valid() error {
	if time.Now().After(j.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}