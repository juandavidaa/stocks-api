package seeds

import (
	"database/sql"
	"fmt"

	sqlrepo "github.com/juandavidaa/stocks-api/internal/infra/persistence/sql"
)

func Seed(db *sql.DB) error {
	fmt.Println("Seeding...")
	if err := SeedStocks(db); err != nil {
		return err
	}
	if err := SeedUsers(sqlrepo.User(db)); err != nil {
		return err
	}
	fmt.Println("Seeded!")
	return nil
}
