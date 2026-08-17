// Package store
package store

import "database/sql"

type Store struct {
	Auth    *AuthStore
	Session *SessionStore
	User    *UserStore
}

func New(db *sql.DB) *Store {
	return &Store{
		Auth: &AuthStore{
			db: db,
		},
		Session: &SessionStore{
			db: db,
		},
		User: &UserStore{
			db: db,
		},
	}
}
