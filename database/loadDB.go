package database

import (
	"encoding/json"
	"os"
)

func (db *DB) LoadDB() (DBStructure, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	file, err := os.ReadFile(db.path)

	if err != nil {
		return DBStructure{}, err
	}

	dbStruct := DBStructure{
		Chirps: make(map[int]Chirp),
		Users:  make(map[int]User),
	}

	if len(file) == 0 {
		return dbStruct, nil
	}

	marshallErr := json.Unmarshal(file, &dbStruct)

	if marshallErr != nil {
		return DBStructure{}, marshallErr
	}

	return dbStruct, nil
}
