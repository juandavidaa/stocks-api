package repository

import (
	"context"

	"github.com/juandavidaa/stocks-api/internal/dto"
)

type BestStockResponse struct {
	Ticker        string  `db:"ticker" json:"ticker"`
	Company       string  `db:"company" json:"company"`
	TargetTo      float32 `db:"target_to" json:"target_to"`
	LastPrice     float32 `db:"last_price" json:"last_price"`
	UpsidePct     float32 `db:"upside_pct" json:"upside_pct"`
	RatingDelta   int     `db:"rating_delta" json:"rating_delta"`
	RecencyWeight float32 `db:"recency_weight" json:"recency_weight"`
	ScoreBase     float32 `db:"score_base" json:"score_base"`
	IsBest        bool    `db:"is_best" json:"is_best"`
}

type StockRepository interface {
	GetBestStocks(ctx context.Context, dto dto.GetStocks) (*[]BestStockResponse, int, error)
}
