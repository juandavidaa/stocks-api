package sql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/juandavidaa/stocks-api/core"
	"github.com/juandavidaa/stocks-api/internal/domain"
	"github.com/juandavidaa/stocks-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type SQLUserRepository struct {
	db *sql.DB
}

func User(db *sql.DB) *SQLUserRepository {
	return &SQLUserRepository{db: db}
}

func (r *SQLUserRepository) Save(ctx context.Context, u *domain.User) error {
	query := `
        INSERT INTO users (name, email, password_hash, created_at)
        VALUES ($1, $2, $3, $4)
		RETURNING id
    `
	fmt.Println("saving user", u.Name, u.Email, u.PasswordHash, u.CreatedAt)
	err := r.db.QueryRowContext(ctx, query, u.Name, u.Email, u.PasswordHash, u.CreatedAt).Scan(&u.ID)

	if err != nil {
		return fmt.Errorf("error saving user: %w", err)
	}

	return nil
}

func (r *SQLUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	row := r.db.QueryRowContext(ctx,
		`SELECT id,name,email,password_hash,created_at
         FROM users WHERE email=$1`, email)

	var u domain.User
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *SQLUserRepository) Login(ctx context.Context, email string, password string) (*repository.Token, error) {
	cfg := core.ConfigInstance()
	secret := cfg.JWTSecret

	var user struct {
		ID           string `db:"id"`
		PasswordHash string `db:"password_hash"`
	}

	query := `SELECT id, password_hash FROM users WHERE email = $1`
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("invalid credentials")
		}
		return nil, fmt.Errorf("database error: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	expiration := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": expiration.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, fmt.Errorf("error generating token: %w", err)
	}

	return &repository.Token{
		Jwt:      tokenString,
		Exp:      expiration.Unix(),
		Jwt_type: "Bearer",
	}, nil
}
