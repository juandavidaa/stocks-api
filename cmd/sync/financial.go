package main

import (
	"context"
	"fmt"
	"os"
	"time"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
	"github.com/juandavidaa/stocks-api/internal/dto"
)

type Quote struct {
	CurrentPrice float64 `json:"c"`
}

var (
	finnhubClient *finnhub.DefaultApiService
)

func Init() {
	apiKey := os.Getenv("FINNHUB_API_KEY")
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", apiKey)
	finnhubClient = finnhub.NewAPIClient(cfg).DefaultApi
}

func GetCurrentPrice(stock *dto.CreateStock) error {
	time.Sleep(1 * time.Second)

	res, _, err := finnhubClient.Quote(context.Background()).Symbol(stock.Ticker).Execute()

	if err != nil {
		return fmt.Errorf("error fetching price for %s: %w", stock.Ticker, err)
	}
	if res.C == nil {
		return fmt.Errorf("price not found for ticker %s", stock.Ticker)
	}
	stock.LastPrice = *res.C

	return err
}
