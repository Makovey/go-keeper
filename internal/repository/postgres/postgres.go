package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/Makovey/go-keeper/internal/config"
	"github.com/Makovey/go-keeper/internal/logger"
)

type Repo struct {
	log logger.Logger
	db  *pgxpool.Pool
}

func NewPostgresRepo(cfg config.Config, log logger.Logger) (*Repo, error) {
	fn := "postgres.NewPostgresRepo"

	pool, err := pgxpool.New(context.Background(), cfg.DatabaseDSN())
	if err != nil {
		return nil, fmt.Errorf("[%s]: %v", fn, err)
	}

	return &Repo{
		log: log,
		db:  pool,
	}, nil
}

func (r *Repo) Close() error {
	r.db.Close()
	return nil
}
