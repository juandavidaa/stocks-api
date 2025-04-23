package repository

import (
	"context"

	"github.com/juandavidaa/stocks-api/internal/domain"
)

type Token struct {
	Jwt      string `json:"jwt"`
	Exp      int64  `json:"exp"`
	Jwt_type string `json:"jwt_type"`
}

type UserRepository interface {
	Save(ctx context.Context, u *domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Login(ctx context.Context, email string, password string) (*Token, error)
}
