package store

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username  string   `json:"username"`
	Password  password `json:"-"`
	CreatedAt string   `json:"created_at"`
}

type password struct {
	text         *string
	hashPassword []byte
}

func (p *password) hash(text string) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	p.text = &text
	p.hashPassword = hashPassword
	return nil
}

func (p *password) Compare(text string) bool {
	return bcrypt.CompareHashAndPassword(p.hashPassword, []byte(text)) == nil
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(username, password string) error {
	user := User{Username: username}
	err := user.Password.hash(password)
	if err != nil {
		return err
	}

	const query = "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err = s.db.Exec(query, user.Username, user.Password.hashPassword)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserStore) GetByUsername(username string) (*User, error) {
	type user struct {
		username string
		password string
	}
	var existingUser user
	const query = "SELECT username, password FROM users WHERE username=?"
	err := s.db.QueryRow(query, username).Scan(&existingUser.username, &existingUser.password)
	if err != nil {
		return nil, err
	}
	password := password{hashPassword: []byte(existingUser.password)}
	return &User{
		Username: existingUser.username,
		Password: password,
	}, nil
}

func (s *UserStore) Verify(username, password string) (bool, error) {
	user, err := s.GetByUsername(username)
	if err != nil {
		return false, err
	}
	return user.Password.Compare(password), nil
}
