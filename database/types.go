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
	Email       string `json:"email"`
	Id          int    `json:"id"`
	Password    []byte `json:"password"`
	IsChirpyRed bool   `json:"is_chirpy_red"`
}

func (u *User) UpdateUserEmail(email string) {
	u.Email = email
}

type UserWithNoPass struct {
	Email       string `json:"email"`
	Id          int    `json:"id"`
	IsChirpyRed bool   `json:"is_chirpy_red"`
}

type DBStructure struct {
	Chirps        []Chirp      `json:"chirps"`
	Users         map[int]User `json:"users"`
	RevokedTokens map[string]time.Time
}
