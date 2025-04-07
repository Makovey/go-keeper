package utils

import (
	"fmt"
	"os"
)

//go:generate mockgen -source=dir_manager.go -destination=mock/dir_manager_mock.go -package=mock
type DirManager interface {
	CreateDir(rootDir, path string) error
	CreateFile(name string) (*os.File, error)
	ReadFile(name string) ([]byte, error)
	RemoveFile(name string) error
}

type dirManager struct {
}

func NewDirManager() DirManager {
	return &dirManager{}
}

func (d *dirManager) CreateDir(rootDir, path string) error {
	fn := "file_storager.createDirIfNeeded"

	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		if err = os.MkdirAll(rootDir, os.ModePerm); err != nil {
			return fmt.Errorf("[%s]: failed to create root dir: %v", fn, err)
		}
	}

	fullPath := fmt.Sprintf("./%s/%s", rootDir, path)

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

func (d *dirManager) CreateFile(name string) (*os.File, error) {
	return os.Create(name)
}

func (d *dirManager) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (d *dirManager) RemoveFile(name string) error {
	return os.Remove(name)
}
