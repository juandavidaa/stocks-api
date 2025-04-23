package main

import (
	"fmt"
	"log"

	"github.com/piquette/finance-go/equity"
)

func main() {
	tickers := []string{"AAPL", "MSFT", "GOOGL", "AMZN", "TSLA"}

	for _, ticker := range tickers {
		eq, err := equity.Get(ticker)
		if err != nil {
			fmt.Printf("Error fetching %s: %v", ticker, err)
			continue
		}
		if eq == nil {
			log.Printf("No data for %s", ticker)
			continue
		}

		fmt.Printf("📈 %s (%s)\n", eq.LongName, eq.Symbol)
		fmt.Printf("  Price: $%.2f\n", eq.RegularMarketPrice)
		fmt.Printf("  Avg Volume 30D: %d\n", eq.AverageDailyVolume3Month)
		fmt.Printf("  Market Cap: $%d\n", eq.MarketCap)
		fmt.Printf("  Shares Outstanding: %d\n", eq.SharesOutstanding)
		fmt.Println()
	}
}
