package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Makovey/go-keeper/internal/repository/entity"
	serviceErrors "github.com/Makovey/go-keeper/internal/service"
)

const (
	errUniqueViolatesCode = "23505"
)

func (r *Repo) RegisterUser(ctx context.Context, user *entity.User) error {
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

func (r *Repo) GetUserInfo(ctx context.Context, email string) (*entity.User, error) {
	fn := "postgres.GetUserInfo"

	row := r.db.QueryRow(
		ctx,
		`SELECT id, name, email, password_hash FROM users WHERE email = $1`,
		email,
	)

	var user entity.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return &entity.User{}, fmt.Errorf("[%s]: %w", fn, serviceErrors.ErrUserNotFound)
		default:
			return &entity.User{}, fmt.Errorf("[%s]: %w", fn, err)
		}
	}

	return &user, nil
}
