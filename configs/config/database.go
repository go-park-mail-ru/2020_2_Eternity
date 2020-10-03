package config

import (
	"fmt"
	"github.com/jackc/pgx"
)

type Database struct {
	config   *ConfDB
	database *pgx.Conn
}

func newDatabase(config *ConfDB) *Database {
	return &Database{
		config: config,
	}
}

func (db *Database) Open() error {
	config, err := pgx.ParseConnectionString(
		fmt.Sprintf(
			"user=%s password=%s dbname=%s sslmode=%s",
			db.config.Postgres.Username,
			db.config.Postgres.Password,
			db.config.Postgres.DbName,
			db.config.Postgres.SslMode,
			))
	if err != nil {
		return err
	}

	db.database, err = pgx.Connect(config)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) Close() error {
	return db.database.Close()
}
