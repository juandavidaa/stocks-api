package users

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/juandavidaa/stocks-api/internal/domain"
	"github.com/juandavidaa/stocks-api/internal/dto"
	"github.com/juandavidaa/stocks-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type CreateResponse struct {
	ID    string
	Name  string
	Email string
}

type Create struct {
	Repo repository.UserRepository
}

var (
	ErrEmailInUse   = errors.New("email already registered")
	ErrInvalidInput = errors.New("invalid user data")
)

func (uc Create) Execute(ctx context.Context, req dto.CreateUser) (CreateResponse, int, error) {
	if req.Name == "" || req.Email == "" || req.Password == "" {
		return CreateResponse{}, http.StatusBadRequest, ErrInvalidInput
	}

	exists, err := uc.Repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return CreateResponse{}, http.StatusInternalServerError, err
	}
	if exists != nil {
		return CreateResponse{}, http.StatusConflict, ErrEmailInUse
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return CreateResponse{}, http.StatusInternalServerError, err
	}

	u := domain.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hash),
		CreatedAt:    time.Now(),
	}

	if err := uc.Repo.Save(ctx, &u); err != nil {
		return CreateResponse{}, http.StatusInternalServerError, err
	}

	return CreateResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}, http.StatusCreated, nil
}
