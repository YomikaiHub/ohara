package store

import (
	"database/sql"
)

type Storage struct {
}

func NewDBStorage(db *sql.DB) Storage {
	return Storage{}
}
