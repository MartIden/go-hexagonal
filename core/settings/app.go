package settings

import (
	"github.com/caarlos0/env/v7"

	dbSettings "github.com/MartIden/go-hexagonal/core/settings/db"
)

type AppSettings struct {
	Port      int    `env:"PORT" envDefault:"3000"`
	BindAddr  string `env:"BIND_ADDR" envDefault:"0.0.0.0:${PORT}" envExpand:"true"`
	Env       string `env:"ENVIRON"`
	Postgres *dbSettings.PostgresSettings
}

func GetAppSettings() (*AppSettings, error) {
	appSettings := AppSettings{}
	pgSettings, pgErr := dbSettings.GetPostrgesSettings()

	if err := env.Parse(&appSettings); err != nil {
		return nil, err
	}

	if pgErr != nil {
		return nil, pgErr
	}

	appSettings.Postgres = pgSettings
	return &appSettings, nil
}
