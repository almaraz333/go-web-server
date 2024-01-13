package database

func (db *DBStructure) CreateChirp(body string, id int, chirpId int) (Chirp, error) {
	successBody := Chirp{
		Body:     body,
		Id:       chirpId,
		AuthorId: id,
	}

	db.Chirps = append(db.Chirps, successBody)

	return successBody, nil
}
