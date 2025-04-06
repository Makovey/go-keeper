package utils

import (
	"fmt"
	"os"
)

func CreateDirIfNeeded(rootDir, path string) error {
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
