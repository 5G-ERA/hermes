package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func EnsurePathExists(path string) error {
	_, dirErr := os.Stat(path)

	if dirErr != nil {
		if os.IsNotExist(dirErr) {
			// The file or directory does not exist
			fmt.Println("File or directory does not exist:", path)

			// Extract the directory part of the path
			dir := filepath.Dir(path)

			// Create the directory if it doesn't exist
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				return err
			}

			fmt.Println("Directory created:", dir)
		} else {
			// Handle other errors
			fmt.Println("Error:", dirErr)
			return dirErr
		}
	} else {
		// The path points to an existing file
		fmt.Println("Path is a file:", path)
	}
	return nil
}
