package core

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB(c *Config) *sql.DB {

	initialDSN := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/defaultdb?sslmode=%s",
		c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.SSLMode,
	)

	initial, err := sql.Open("postgres", initialDSN)
	if err != nil {
		panic(err)
	}

	initial.Exec("CREATE DATABASE IF NOT EXISTS " + c.DBName)

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName, c.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(fmt.Sprintf("failed to open DB: %s", err))
	}
	return db
}
