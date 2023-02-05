package cmd

import (
	"net/http"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx"

	"github.com/MartIden/go-hexagonal/core/databases"
	corePostgres "github.com/MartIden/go-hexagonal/core/databases/postgres"
	"github.com/MartIden/go-hexagonal/core/server"
	"github.com/MartIden/go-hexagonal/core/settings"
	"github.com/MartIden/go-hexagonal/core/store/sqlstore"
)

func Start(settings *settings.AppSettings) error {
	connector := corePostgres.GetPostgresConnector(settings)
	db, err := databases.GetDB(connector, settings)
	
	if err != nil {
		return err
	}

	//Migrations
	// driver, err := postgres.WithInstance(db, &postgres.Config{})
	// m, err := migrate.NewWithDatabaseInstance(
    //     "file:///migrations/postgres",
    //     connector.GetDriverName(), driver)
    // m.Up()

	store := sqlstore.New(db)
	srv := server.NewServer(*store)
	return http.ListenAndServe(settings.BindAddr, srv)
}