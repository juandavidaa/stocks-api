package users

import (
	"database/sql"

	sqlrepo "github.com/juandavidaa/stocks-api/internal/infra/persistence/sql"
	userUC "github.com/juandavidaa/stocks-api/internal/usecases/users"
)

func New(db *sql.DB) Module {
	repo := sqlrepo.User(db)

	handlers := Handlers{
		CreateUC: userUC.Create{Repo: repo},
		LoginUC:  userUC.Login{Repo: repo},
	}
	return Module{Handlers: handlers}
}
