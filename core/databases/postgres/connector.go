package postgres

import (
	"fmt"

	"github.com/MartIden/go-hexagonal/core/settings"
)

type PostgresConnector struct {
	driverName string
	settings *settings.AppSettings
}

func (pgc *PostgresConnector) GetDSN() string {
	pgSettings := pgc.settings.Postgress
	
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
    	pgSettings.Host, 
		pgSettings.Port, 
		pgSettings.Login, 
		pgSettings.Password, 
		pgSettings.DbName, 
		pgSettings.SslMode,
	)
	
	return psqlInfo
}

func (pgc *PostgresConnector) GetDriverName() string {
	return pgc.driverName
}

func GetPostgresConnector(settings *settings.AppSettings) *PostgresConnector {
	return &PostgresConnector{
		driverName: "postgres",
		settings: settings,
	}
}