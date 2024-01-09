package database

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (db *DB) UpdateUser(userId int, email string, dbStructure DBStructure, password string) (User, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 4)

	if err != nil {
		return User{}, errors.New("Could not has password for new user")
	}

	newUser := dbStructure.Users[userId]
	newUser.Email = email
	newUser.Password = hashedPass

	dbStructure.Users[userId] = newUser

	return newUser, nil
}
