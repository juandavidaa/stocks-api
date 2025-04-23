package users

import (
	"context"
	"errors"
	"net/http"

	"github.com/juandavidaa/stocks-api/internal/repository"
)

type LoginRequest struct {
	Email    string
	Password string
}

type Login struct {
	Repo   repository.UserRepository
	Secret string
}

var ErrBadCredentials = errors.New("invalid email or password")

func (uc Login) Execute(ctx context.Context, req LoginRequest) (repository.Token, int, error) {
	u, err := uc.Repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return repository.Token{}, http.StatusInternalServerError, err
	}
	if u == nil {
		return repository.Token{}, http.StatusUnauthorized, ErrBadCredentials
	}

	response, error := uc.Repo.Login(ctx, req.Email, req.Password)

	if error != nil {
		return repository.Token{}, http.StatusInternalServerError, error
	}

	return *response, http.StatusOK, nil
}
