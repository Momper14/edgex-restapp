package util

import (
	"log"
	"os"
	"path/filepath"
)

// CreateParents creates all parent directories
func CreateParents(file string) {
	if err := os.MkdirAll(filepath.Dir(file), 0700); err != nil {
		log.Fatal(err)
	}
}

// FileExists checks if the file exists
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}
