package util

import (
	"os"
	"path/filepath"
)

func EnsurePathToFileExists(path string) error {
	dir := filepath.Dir(path)
	return os.MkdirAll(dir, 0755)
}

// ReadAllFiles lists all files form the directory including subdirectories
func ReadAllFiles(root string) ([]os.DirEntry, error) {
	var files []os.DirEntry
	err := filepath.WalkDir(root, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, info)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
