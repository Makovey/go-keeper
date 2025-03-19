package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Makovey/go-keeper/internal/repository/dbo"
	serviceErrors "github.com/Makovey/go-keeper/internal/service"
)

const (
	errUniqueViolatesCode = "23505"
)

func (r *Repo) RegisterUser(ctx context.Context, user *dbo.User) error {
	fn := "postgres.RegisterUser"

	_, err := r.db.Exec(
		ctx,
		`INSERT INTO users (id, name, email, password_hash) VALUES ($1, $2, $3, $4)`,
		user.ID,
		user.Name,
		user.Email,
		user.PasswordHash,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok && pgErr.Code == errUniqueViolatesCode {
			return fmt.Errorf("[%s]: %w", fn, serviceErrors.ErrUserAlreadyExists)
		}

		return fmt.Errorf("[%s]: %w", fn, err)
	}

	return nil
}
