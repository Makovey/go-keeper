package file_storager

import (
	"bufio"
	"fmt"
	"io"
	"sync"

	"github.com/Makovey/go-keeper/internal/logger"
	"github.com/Makovey/go-keeper/internal/service/storage"
	"github.com/Makovey/go-keeper/internal/utils"
)

const rootDirForStorage = "file_storage"

type diskStorage struct {
	log        logger.Logger
	dirManager utils.DirManager
	mu         sync.RWMutex
}

func NewDiskStorage(
	log logger.Logger,
	dirManager utils.DirManager,
) storage.FileStorager {
	return &diskStorage{
		log:        log,
		dirManager: dirManager,
		mu:         sync.RWMutex{},
	}
}

func (d *diskStorage) Save(path, fileName string, data *bufio.Reader) error {
	fn := "file_storager.Save"
	d.mu.Lock()
	defer d.mu.Unlock()

	if err := d.dirManager.CreateDir(rootDirForStorage, path); err != nil {
		return fmt.Errorf("[%s]: %v", fn, err)
	}

	fullPath := fmt.Sprintf("./%s/%s/%s", rootDirForStorage, path, fileName)
	file, err := d.dirManager.CreateFile(fullPath)
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

func (d *diskStorage) Get(path string) ([]byte, error) {
	fn := "file_storager.Get"

	d.mu.RLock()
	defer d.mu.RUnlock()

	fullPath := fmt.Sprintf("./%s/%s", rootDirForStorage, path)
	data, err := d.dirManager.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("[%s]: can't read file: %v", fn, err)
	}

	return data, nil
}

func (d *diskStorage) Delete(path string) error {
	fn := "file_storager.Delete"
	d.mu.Lock()
	defer d.mu.Unlock()

	fullPath := fmt.Sprintf("./%s/%s", rootDirForStorage, path)
	if err := d.dirManager.RemoveFile(fullPath); err != nil {
		return fmt.Errorf("[%s]: can't delete file: %v", fn, err)
	}

	return nil
}
