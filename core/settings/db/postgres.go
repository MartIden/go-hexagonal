package db

import (
	"github.com/caarlos0/env/v7"
)

type PostgresSettings struct {
	Host     string `env:"PG_HOST"`
	Port     string `env:"PG_PORT"`
	Login    string `env:"PG_LOGIN"`
	Password string `env:"PG_PASSWORD"`
	DbName   string `env:"PG_DBNAME"`
	SslMode  string `env:"PG_SSL_mode" envDefault:"disable"`
}

func (pgs *PostgresSettings) InitFromEnv() error {
	if err := env.Parse(pgs); err != nil {
		return err
	}
	
	return nil
}

func GetPostrgesSettings() (*PostgresSettings, error) {
	pgSettings := PostgresSettings{}
	pgErr := pgSettings.InitFromEnv()

	if pgErr != nil {
		return nil, pgErr
	}

	return &pgSettings, nil
}