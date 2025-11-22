package store

import (
	"database/sql"
)

type UserStorer interface {
	Create(username, password string) error
	Verify(username, password string) (bool, error)
	GetByUsername(username string) (*User, error)
}

type Storage struct {
	Users UserStorer
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users: &UserStore{db},
	}
}
