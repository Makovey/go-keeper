package entity

import "time"

type File struct {
	ID          string
	OwnerID     string
	FileName    string
	FileSize    int
	Path        string
	IsEncrypted bool
	CreatedAt   time.Time
}
