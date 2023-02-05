package cmd

import (
	"log"
	"net/http"

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
		log.Fatalln(err.Error())
		return err
	}

	migration := corePostgres.CreatePostgresMigration(db)
	
	if err := migration.Down(); err != nil {
		log.Fatalln(err.Error())
		return err
	}
		
	store := sqlstore.New(db)
	srv := server.NewServer(*store)
	return http.ListenAndServe(settings.BindAddr, srv)
}