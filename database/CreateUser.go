package database

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (db *DBStructure) CreateUser(email string, id int, password string) (UserWithNoPass, error) {
	_, ok := db.Users[id]

	if ok {
		return UserWithNoPass{}, errors.New("User already exists")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 4)

	if err != nil {
		return UserWithNoPass{}, errors.New("Could not has password for new user")
	}

	newUser := User{
		Email:    email,
		Id:       id,
		Password: hashedPass,
	}

	db.Users[id] = newUser

	return UserWithNoPass{Id: newUser.Id, Email: newUser.Email}, nil
}
