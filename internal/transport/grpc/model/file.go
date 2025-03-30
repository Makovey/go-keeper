package model

import "bytes"

type File struct {
	Data     bytes.Reader
	FileName string
	FileSize int
}
