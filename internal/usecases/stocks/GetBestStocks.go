package stocks

import (
	"context"

	"github.com/juandavidaa/stocks-api/internal/dto"

	"github.com/juandavidaa/stocks-api/internal/repository"
)

type GetBestStocks struct {
	Repo repository.StockRepository
}

func (u GetBestStocks) Execute(ctx context.Context, req dto.GetStocks) (*[]repository.BestStockResponse, int, error) {

	stocks, status, err := u.Repo.GetBestStocks(ctx, req)
	if err != nil {
		return nil, status, err
	}
	return stocks, status, nil
}
