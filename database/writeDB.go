package database

import (
	"encoding/json"
	"os"
)

func (db *DB) WriteDB(DBStructure DBStructure) error {
	bytes, err := json.Marshal(DBStructure)

	if err != nil {
		return err
	}

	writeErr := os.WriteFile(db.path, bytes, 0664)

	if writeErr != nil {
		return err
	}

	return nil
}
