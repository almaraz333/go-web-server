package database

func (db *DBStructure) GetChirps() ([]Chirp, error) {
	var chirps []Chirp

	for _, chirp := range db.Chirps {
		chirps = append(chirps, chirp)
	}

	return chirps, nil
}
