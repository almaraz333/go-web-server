package database

import "sync"

type DB struct {
	path string
	mu   *sync.RWMutex
}

type Chirp struct {
	Id   int    `json:"id"`
	Body string `json:"body"`
}

type User struct {
	Email    string `json:"email"`
	Id       int    `json:"id"`
	Password []byte `json:"password"`
}

func (u *User) UpdateUserEmail(email string) {
	u.Email = email
}

type UserWithNoPass struct {
	Email string `json:"email"`
	Id    int    `json:"id"`
}

type DBStructure struct {
	Chirps map[int]Chirp `json:"chirps"`
	Users  map[int]User  `json:"users"`
}
