package model

import (
	"bufio"
	"time"
)

type File struct {
	Data     bufio.Reader
	FileName string
	FileSize int
}

type ExtendedInfoFile struct {
	ID        string
	FileName  string
	FileSize  string
	CreatedAt time.Time
}
