package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/Makovey/go-keeper/internal/config"
	"github.com/Makovey/go-keeper/internal/logger"
)

type Repo struct {
	log logger.Logger
	db  *sql.DB
}

func NewPostgresRepo(cfg config.Config, log logger.Logger) (*Repo, error) {
	fn := "postgres.NewPostgresRepo"

	db, err := sql.Open("pgx", cfg.DatabaseDSN())
	if err != nil {
		return nil, fmt.Errorf("[%s]: %v", fn, err)
	}

	return &Repo{
		log: log,
		db:  db,
	}, nil
}
