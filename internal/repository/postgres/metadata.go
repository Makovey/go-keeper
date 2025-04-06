package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/Makovey/go-keeper/internal/repository/entity"
	serviceErrors "github.com/Makovey/go-keeper/internal/service"
)

func (r *Repo) SaveFileMetadata(ctx context.Context, fileData *entity.File) error {
	fn := "postgres.SaveFileMetadata"

	_, err := r.db.Exec(
		ctx,
		`INSERT INTO files_metadata (
				id, 
				owner_user_id, 
				file_name, 
				file_size, 
				path
		) VALUES ($1, $2, $3, $4, $5)`,
		fileData.ID,
		fileData.OwnerID,
		fileData.FileName,
		fileData.FileSize,
		fileData.Path,
	)

	if err != nil {
		return fmt.Errorf("[%s]: %w", fn, err)
	}

	return nil
}

func (r *Repo) GetFileMetadata(ctx context.Context, userID, fileID string) (*entity.File, error) {
	fn := "postgres.GetFileMetadata"

	row := r.db.QueryRow(
		ctx,
		`SELECT * FROM files_metadata WHERE owner_user_id = $1 AND id = $2`,
		userID,
		fileID,
	)

	var file entity.File
	err := row.Scan(&file.ID, &file.OwnerID, &file.FileName, &file.FileSize, &file.Path, &file.CreatedAt)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return &entity.File{}, fmt.Errorf("[%s]: %w", fn, serviceErrors.ErrFileNotFound)
		default:
			return &entity.File{}, fmt.Errorf("[%s]: %w", fn, err)
		}
	}

	return &file, nil
}
