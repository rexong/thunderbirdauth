package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func newDatabase(driverName, addr string) (*sql.DB, error) {
	db, err := sql.Open(driverName, addr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func NewSqlite(addr string) (*sql.DB, error) {
	return newDatabase("sqlite3", addr)
}
