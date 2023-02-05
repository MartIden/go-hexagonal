package postgres

import (
	"github.com/MartIden/go-hexagonal/core/settings"
)

type PostgresConnector struct {
	driverName string
	settings *settings.AppSettings
}

func (pgc *PostgresConnector) GetDSN() string {
	// postgresql://
	// postgresql://localhost
	// postgresql://localhost:5432
	// postgresql://localhost/mydb
	// postgresql://user@localhost
	// postgresql://user:secret@localhost
	// postgresql://other@localhost/otherdb?connect_timeout=10&application_name=myapp
	// postgresql://localhost/mydb?user=other&password=secret
	// pgSettings := pgc.settings.Postgres
	
	// psqlInfo := fmt.Sprintf(
	// 	"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
    // 	pgSettings.Host, 
	// 	pgSettings.Port, 
	// 	pgSettings.Login, 
	// 	pgSettings.Password, 
	// 	pgSettings.DbName, 
	// 	pgSettings.SslMode,
	// )
	
	// return psqlInfo

	return "postgres://golang:golang@127.0.0.1:5437/webhooks?sslmode=disable"
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