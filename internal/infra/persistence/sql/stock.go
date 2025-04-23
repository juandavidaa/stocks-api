package sql

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/juandavidaa/stocks-api/internal/dto"
	"github.com/juandavidaa/stocks-api/internal/repository"
)

type StockSqlRepository struct {
	db *sql.DB
}

func Stock(db *sql.DB) *StockSqlRepository {
	return &StockSqlRepository{db: db}
}

func (r StockSqlRepository) GetBestStocks(ctx context.Context, dto dto.GetStocks) (*[]repository.BestStockResponse, int, error) {

	var whereClauses []string
	switch dto.Risk {
	case "high":
		whereClauses = append(whereClauses, `last_price >= 1
                 AND upside_pct   > 0
                 AND score_base   > 0`)
	case "medium":
		whereClauses = append(whereClauses, `last_price >= 1
                 AND upside_pct   > 0
                 AND score_base   > 0
                 AND (target_to - target_from) >= 0`)
	case "low":
		whereClauses = append(whereClauses, `last_price >= 1
                 AND upside_pct   > 0
                 AND score_base   > 0
                 AND (target_to - target_from) > 2`)
	default:
		return nil, http.StatusBadRequest, fmt.Errorf("invalid risk level: %q", dto.Risk)
	}

	if dto.Query != "" {
		whereClauses = append(whereClauses, `(ticker ILIKE $2 OR company ILIKE $2)`)
	}

	where := strings.Join(whereClauses, " AND ")

	query := fmt.Sprintf(`
        SELECT
          ticker,
          company,
          target_to,
          last_price,
          upside_pct,
          rating_delta,
          recency_weight,
          score_base
        FROM stocks
        WHERE %s
        ORDER BY score_base DESC
        LIMIT 20
        OFFSET $1
    `, where)

	args := []interface{}{
		dto.Page * 20,
	}
	if dto.Query != "" {
		args = append(args, "%"+dto.Query+"%")
	}

	rows, err := r.db.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer rows.Close()

	var results []repository.BestStockResponse
	index := 0
	for rows.Next() {
		var s repository.BestStockResponse
		err := rows.Scan(
			&s.Ticker,
			&s.Company,
			&s.TargetTo,
			&s.LastPrice,
			&s.UpsidePct,
			&s.RatingDelta,
			&s.RecencyWeight,
			&s.ScoreBase,
		)

		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		s.IsBest = index == 0 && dto.Query == ""
		results = append(results, s)

		index++
	}

	if err := rows.Err(); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &results, http.StatusOK, nil
}
