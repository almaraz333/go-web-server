package database

import (
	"errors"
	"strconv"
)

func (db *DBStructure) GetChirpByID(id int) (Chirp, error) {
	for _, chirp := range db.Chirps {
		if chirp.Id == id {
			return chirp, nil
		}
	}

	return Chirp{}, errors.New("Cannot find chirp with ID: " + strconv.Itoa(id))
}
