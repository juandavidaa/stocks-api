package main

import (
	"github.com/juandavidaa/stocks-api/core"
	"github.com/juandavidaa/stocks-api/internal/migrations"
	"github.com/juandavidaa/stocks-api/internal/seeds"
)

func main() {
	cfg := core.ConfigInstance()
	db := core.ConnectDB(cfg)
	defer db.Close()

	if err := migrations.Migrate(db); err != nil {
		panic(err)
	}

	if err := seeds.Seed(db); err != nil {
		panic(err)
	}
}
