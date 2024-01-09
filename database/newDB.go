package database

import (
	"errors"
	"os"
	"path/filepath"
	"sync"
)

func NewDB(path string) (*DB, error) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		_, err := os.Create(path)
		if err != nil {
			return nil, errors.New("Could not create DB...")
		}
	}

	filePath, err := filepath.Abs(path)

	if err != nil {
		return nil, errors.New("Could not resolve file path")
	}

	createdDB := DB{
		path: filePath,
		mu:   &sync.RWMutex{},
	}

	return &createdDB, nil
}
