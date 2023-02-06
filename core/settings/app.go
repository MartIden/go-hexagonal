package settings

import (
	"github.com/caarlos0/env/v7"

	dbSettings "github.com/MartIden/go-hexagonal/core/settings/db"
)

type AppSettings struct {
	Port     int    `env:"PORT" envDefault:"3000"`
	BindAddr string `env:"BIND_ADDR" envDefault:"0.0.0.0:${PORT}" envExpand:"true"`
	Env      string `env:"ENVIRON"`
	Postgres *dbSettings.PostgresSettings
}

func includePostrgesSettings(appSettings *AppSettings) error {
	if pgSettings, err := dbSettings.GetPostrgesSettings(); err == nil {
		appSettings.Postgres = pgSettings
		return nil
	} else {
		return err
	}
}

func GetAppSettings() (*AppSettings, error) {
	appSettings := AppSettings{}

	if err := env.Parse(&appSettings); err != nil {
		return nil, err
	}

	if err := includePostrgesSettings(&appSettings); err != nil {
		return nil, err
	}

	return &appSettings, nil
}
