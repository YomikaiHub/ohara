// Package store
package store

import "database/sql"

type Store struct{}

func New(db *sql.DB) *Store {
	return &Store{}
}
