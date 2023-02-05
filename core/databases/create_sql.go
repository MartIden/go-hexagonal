package databases

import (
	"database/sql"

	"github.com/MartIden/go-hexagonal/core/settings"

	_ "github.com/jackc/pgx"
)

type IDBConnectable interface {
	GetDSN() string
	GetDriverName() string
}

func GetDB(
	connector IDBConnectable,
	settings *settings.AppSettings,
) (*sql.DB, error) {

	db, err := sql.Open(
		connector.GetDriverName(),
		connector.GetDSN(),
	)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
