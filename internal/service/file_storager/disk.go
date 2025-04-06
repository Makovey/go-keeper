package file_storager

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/Makovey/go-keeper/internal/logger"
	"github.com/Makovey/go-keeper/internal/service/storage"
)

const rootDirForStorage = "fileStorage"

type diskStorager struct {
	log logger.Logger
	mu  sync.RWMutex
}

func NewDiskStorager(
	log logger.Logger,
) storage.FileStorager {
	return &diskStorager{
		log: log,
		mu:  sync.RWMutex{},
	}
}

func (d *diskStorager) Save(path, fileName string, data *bufio.Reader) error {
	fn := "file_storager.Save"
	d.mu.Lock()
	defer d.mu.Unlock()

	if err := createDirIfNeeded(path); err != nil {
		return fmt.Errorf("[%s]: %v", fn, err)
	}

	fullPath := fmt.Sprintf("./%s/%s/%s", rootDirForStorage, path, fileName)
	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("[%s]: failed to create file: %v", fn, err)
	}
	defer file.Close()

	bufWriter := bufio.NewWriter(file)
	defer bufWriter.Flush()

	_, err = io.Copy(bufWriter, data)
	if err != nil {
		return fmt.Errorf("[%s]: failed to write data: %v", fn, err)
	}

	return nil
}

func (d *diskStorager) Get(path string, size int) (*bufio.Reader, error) {
	fn := "file_storager.Get"

	d.mu.RLock()
	defer d.mu.RUnlock()

	fullPath := fmt.Sprintf("./%s/%s", rootDirForStorage, path)
	file, err := os.Open(fullPath)
	if err != nil {
		return nil, fmt.Errorf("[%s]: can't open file: %v", fn, err)
	}
	defer file.Close()

	return bufio.NewReaderSize(file, size), nil
}

func createDirIfNeeded(path string) error {
	fn := "file_storager.createDirIfNeeded"

	if _, err := os.Stat(rootDirForStorage); os.IsNotExist(err) {
		if err = os.MkdirAll(rootDirForStorage, os.ModePerm); err != nil {
			return fmt.Errorf("[%s]: failed to create root dir: %v", fn, err)
		}
	}

	fullPath := fmt.Sprintf("./%s/%s", rootDirForStorage, path)

	if _, err := os.Stat(fullPath); err != nil {
		if os.IsNotExist(err) {
			if err = os.Mkdir(fullPath, os.ModePerm); err != nil {
				return fmt.Errorf("[%s]: can't create new directory %v", fn, err)
			}
		} else {
			return fmt.Errorf("[%s]: %v", fn, err)
		}
	}

	return nil
}
