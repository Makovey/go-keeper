package postgres

import (
	"context"
	"fmt"

	"github.com/Makovey/go-keeper/internal/repository/entity"
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
