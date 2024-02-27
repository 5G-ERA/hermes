package util

import (
	"os"
	"path/filepath"
)

func EnsurePathToFileExists(path string) error {
	dir := filepath.Dir(path)
	return os.MkdirAll(dir, 0755)
}
