package database

func (db *DBStructure) CreateChirp(body string, id int) (Chirp, error) {
	successBody := Chirp{
		Body: body,
		Id:   id,
	}

	db.Chirps[id] = successBody

	return successBody, nil
}
