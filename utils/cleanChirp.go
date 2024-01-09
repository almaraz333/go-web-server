package utils

import "strings"

func CleanChirp(chirp string) string {
	BAD_WORDS := []string{"kerfuffle", "sharbert", "fornax"}

	splitChirp := strings.Split(chirp, " ")

	for index, word := range splitChirp {
		for _, badWord := range BAD_WORDS {
			if strings.ToLower(word) == badWord {
				splitChirp[index] = "***"
			}
		}
	}

	return strings.Join(splitChirp, " ")
}
