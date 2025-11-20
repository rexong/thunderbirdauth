package store

import (
	"database/sql"
)

type Storage struct {
	Users interface {
		Create(username, password string) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users: &UserStore{db},
	}
}
