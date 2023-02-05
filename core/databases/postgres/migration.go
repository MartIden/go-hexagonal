package postgres

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	postgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type PostgresMigration struct {
	db *sql.DB
}

func (pm *PostgresMigration) createInstance() (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(pm.db, &postgres.Config{})

	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}
    
	m, err := migrate.NewWithDatabaseInstance(
        "file://migrations/postgres",
        "postgres", 
		driver,
	)
	
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	return m, nil
}

func (pm *PostgresMigration) Up() error {
	if migration, err := pm.createInstance(); err == nil{
		return migration.Up()
	} else {
		return err
	}
}

func (pm *PostgresMigration) Down() error {
	if migration, err := pm.createInstance(); err == nil{
		return migration.Down()
	} else {
		return err
	}
}

func CreatePostgresMigration(db *sql.DB) *PostgresMigration {
	return &PostgresMigration{
		db: db,
	}
}
