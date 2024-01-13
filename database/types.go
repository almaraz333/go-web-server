package database

import (
	"sync"
	"time"
)

type DB struct {
	path string
	mu   *sync.RWMutex
}

type Chirp struct {
	Id       int    `json:"id"`
	Body     string `json:"body"`
	AuthorId int    `json:"author_id"`
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
	Chirps        []Chirp      `json:"chirps"`
	Users         map[int]User `json:"users"`
	RevokedTokens map[string]time.Time
}
