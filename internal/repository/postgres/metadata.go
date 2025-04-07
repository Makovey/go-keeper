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

func (r *Repo) GetUsersFiles(ctx context.Context, userID string) ([]*entity.File, error) {
	fn := "postgres.GetUsersFiles"

	rows, err := r.db.Query(
		ctx,
		`SELECT id, file_name, file_size, created_at FROM files_metadata WHERE owner_user_id = $1`,
		userID,
	)

	if err != nil {
		return nil, fmt.Errorf("[%s]: %w", fn, err)
	}
	defer rows.Close()

	var files []*entity.File
	for rows.Next() {
		var file entity.File
		err = rows.Scan(&file.ID, &file.FileName, &file.FileSize, &file.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("[%s]: %w", fn, err)
		}
		files = append(files, &file)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("[%s]: %w", fn, err)
	}

	return files, nil
}

func (r *Repo) DeleteUsersFile(ctx context.Context, userID, fileID string) error {
	fn := "postgres.DeleteUsersFile"

	res, err := r.db.Exec(
		ctx,
		`DELETE FROM files_metadata WHERE owner_user_id = $1 AND id = $2`,
		userID,
		fileID,
	)
	if err != nil {
		return fmt.Errorf("[%s]: %w", fn, err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("[%s]: %w", fn, serviceErrors.ErrFileNotFound)
	}

	return nil
}
