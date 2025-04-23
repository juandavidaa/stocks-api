package migrations

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func MigrateStocks(db *sql.DB) error {

	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS stocks (
		id SERIAL PRIMARY KEY,
		ticker TEXT UNIQUE NOT NULL,
		company TEXT NOT NULL,
		brokerage TEXT NOT NULL,
		action TEXT NOT NULL,
		rating_from TEXT,
		rating_to TEXT,
		target_from DECIMAL,
		target_to DECIMAL,
		last_price DECIMAL,
		upside_pct     DECIMAL,
		rating_delta   SMALLINT,
		recency_weight DECIMAL,
		score_base     DECIMAL,
		event_time TIMESTAMPTZ
	);`)

	fmt.Println("Stocks table created!")
	return err
}
