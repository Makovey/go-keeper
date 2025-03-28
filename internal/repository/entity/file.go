package entity

import "time"

type File struct {
	ID         string
	OwnerID    string
	FileName   string
	FileSize   string
	Path       string
	UploadedAt time.Time
}
