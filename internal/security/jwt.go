package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/juandavidaa/stocks-api/core"
)

func GenerateToken(userID, email string) (string, error) {
	cfg := core.ConfigInstance()
	secret := cfg.JWTSecret
	claims := jwt.MapClaims{
		"sub":   userID,
		"email": email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
		"iat":   time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseToken(tokenStr string) (*jwt.Token, error) {
	cfg := core.ConfigInstance()
	secret := cfg.JWTSecret
	return jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
}
