package sqlstore

import (
	"database/sql"

	p "github.com/MartIden/go-hexagonal/ports/out/persistence"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func User() p.IUserRepoPort {
	return UserRepo{}
}
