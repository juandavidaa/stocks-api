package migrations

import (
	"database/sql"
	"fmt"
)

func Migrate(db *sql.DB) error {

	fmt.Println("Migrating...")

	if err := MigrateUsers(db); err != nil {
		return err
	}
	if err := MigrateStocks(db); err != nil {
		return err
	}
	fmt.Println("Migrated!")
	return nil
}
