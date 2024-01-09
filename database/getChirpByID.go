package database

import (
	"errors"
	"strconv"
)

func (db *DBStructure) GetChirpByID(id int) (Chirp, error) {
	for chirpID, chirp := range db.Chirps {
		if chirpID == id {
			return chirp, nil
		}
	}

	return Chirp{}, errors.New("Cannot find chirp with ID: " + strconv.Itoa(id))
}
