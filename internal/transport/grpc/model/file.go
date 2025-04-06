package model

import (
	"bufio"
)

type File struct {
	Data     bufio.Reader
	FileName string
	FileSize int
}
