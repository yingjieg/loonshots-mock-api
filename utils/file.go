package utils

import (
	"os"
	"path/filepath"
)

func ListFolder(folderPath string) ([]string, error) {
	var files []string

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != ".csv" {
			return nil
		}

		files = append(files, path)

		return nil
	})

	if err != nil {
		return files, err
	}

	return files, nil
}
