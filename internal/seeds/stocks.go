package seeds

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/juandavidaa/stocks-api/internal/dto"
)

func ratingDelta(action string) int {
	if strings.Contains(action, "upgraded by") || strings.Contains(action, "target raised") {
		return 1
	}
	if strings.Contains(action, "downgraded by") || strings.Contains(action, "target lowered") {
		return -1
	}
	return 0
}

func recencyWeight(event time.Time) float32 {
	days := int(time.Since(event).Hours() / 24)
	return 1 / float32(1+days)
}

func SeedStocks(db *sql.DB) error {
	//verify existing data
	if verifyExistingStocks(db) {
		return nil
	}

	file, err := os.Open("seed/data.json")
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	batchSize := 10
	var batch []dto.CreateStock

	for scanner.Scan() {
		var item dto.CreateStock
		if err := json.Unmarshal(scanner.Bytes(), &item); err != nil {
			return fmt.Errorf("error parsing line: %w", err)
		}

		tt, _ := strconv.ParseFloat(item.TargetTo, 32)

		et, _ := time.Parse(time.RFC3339Nano, item.Time)

		item.RatingDelta = ratingDelta(item.Action)
		item.RecencyWeight = recencyWeight(et)

		if item.LastPrice <= 0 {
			item.UpsidePct = 0
			item.ScoreBase = 0

		} else {
			item.UpsidePct = (float32(tt) / item.LastPrice) - 1
			item.ScoreBase = item.UpsidePct * float32(item.RatingDelta+1) * item.RecencyWeight
		}

		batch = append(batch, item)

		if len(batch) >= batchSize {
			if err := insertBatch(db, batch); err != nil {
				return err
			}
			batch = []dto.CreateStock{}
		}
	}

	if len(batch) > 0 {
		if err := insertBatch(db, batch); err != nil {
			return err
		}
	}
	fmt.Println("Stocks seeded!")
	return nil
}

func verifyExistingStocks(db *sql.DB) bool {
	var count int
	db.QueryRow("SELECT COUNT(*) FROM stocks").Scan(&count)
	if count > 0 {
		fmt.Println("⚠️ Skipping stocks seed - Data already exists")
		return true
	}
	return false
}

func insertBatch(db *sql.DB, batch []dto.CreateStock) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt := `
    INSERT INTO stocks (
      ticker, company, brokerage, action,
      rating_from, rating_to, target_from, target_to,
      event_time, last_price,
      upside_pct, rating_delta, recency_weight, score_base
    ) VALUES
    `
	var placeholders []string
	var args []interface{}

	for i, s := range batch {
		off := i * 14
		placeholders = append(placeholders, fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d)",
			off+1, off+2, off+3, off+4, off+5, off+6, off+7,
			off+8, off+9, off+10, off+11, off+12, off+13, off+14))

		args = append(args,
			s.Ticker, s.Company, s.Brokerage, s.Action,
			s.RatingFrom, s.RatingTo, s.TargetFrom, s.TargetTo,
			s.Time, s.LastPrice,
			s.UpsidePct, s.RatingDelta, s.RecencyWeight, s.ScoreBase,
		)
	}

	stmt += strings.Join(placeholders, ", ")
	_, err = tx.Exec(stmt, args...)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("batch insert failed: %w", err)
	}

	return tx.Commit()
}
