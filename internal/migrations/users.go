package migrations

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func MigrateUsers(db *sql.DB) error {

	_, err := db.Exec(`
	  CREATE TABLE IF NOT EXISTS users (
	    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	    name TEXT NOT NULL,
	    email TEXT UNIQUE NOT NULL,
	    password_hash TEXT NOT NULL,
	    created_at TIMESTAMPTZ DEFAULT now()
	  );`)

	fmt.Println("Users table created!")
	return err
}
