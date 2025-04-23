package seeds

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/juandavidaa/stocks-api/internal/domain"
	"github.com/juandavidaa/stocks-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

func SeedUsers(repo repository.UserRepository) error {
	//verify existing data
	if verifyExistingUsers(repo) {
		return nil
	}

	name := os.Getenv("DUMMY_USER")
	email := os.Getenv("DUMMY_EMAIL")
	password := os.Getenv("DUMMY_PASS")

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}

	user := &domain.User{
		Name:         name,
		Email:        email,
		PasswordHash: string(hash),
		CreatedAt:    time.Now(),
	}

	if err := repo.Save(context.Background(), user); err != nil {
		return fmt.Errorf("error saving user %s: %w", email, err)
	}

	fmt.Println("Users seeded!")
	return nil
}

func verifyExistingUsers(repo repository.UserRepository) bool {
	user, _ := repo.FindByEmail(context.Background(), "admin@example.com")
	if user != nil {
		fmt.Println("⚠️ Skipping users seed - Data already exists")
		return true
	}
	return false
}
