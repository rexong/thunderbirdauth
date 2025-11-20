package store

import (
	"database/sql"
)

type Storage struct {
	Users interface {
		Create(username, password string) error
		Verify(username, password string) (bool, error)
		GetByUsername(username string) (*User, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users: &UserStore{db},
	}
}
