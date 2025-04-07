package file_storager

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/Makovey/go-keeper/internal/logger"
	"github.com/Makovey/go-keeper/internal/service/storage"
	"github.com/Makovey/go-keeper/internal/utils"
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

	if err := utils.CreateDirIfNeeded(rootDirForStorage, path); err != nil {
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

func (d *diskStorager) Get(path string) ([]byte, error) {
	fn := "file_storager.Get"

	d.mu.RLock()
	defer d.mu.RUnlock()

	fullPath := fmt.Sprintf("./%s/%s", rootDirForStorage, path)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("[%s]: can't read file: %v", fn, err)
	}

	return data, nil
}

func (d *diskStorager) Delete(path string) error {
	fn := "file_storager.Delete"
	d.mu.Lock()
	defer d.mu.Unlock()

	fullPath := fmt.Sprintf("./%s/%s", rootDirForStorage, path)
	if err := os.Remove(fullPath); err != nil {
		return fmt.Errorf("[%s]: can't delete file: %v", fn, err)
	}

	return nil
}
