package stocks

import (
	"database/sql"

	sqlrepo "github.com/juandavidaa/stocks-api/internal/infra/persistence/sql"
	stockUC "github.com/juandavidaa/stocks-api/internal/usecases/stocks"
)

func New(db *sql.DB) Module {
	repo := sqlrepo.Stock(db)

	handlers := Handlers{
		GetBestStocksUC: stockUC.GetBestStocks{Repo: repo},
	}
	return Module{Handlers: handlers}
}
