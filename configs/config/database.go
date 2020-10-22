package config

import (
	"fmt"
	"github.com/jackc/pgx"
)

type Database struct {
	config *ConfDB
}

func newDatabase(config *ConfDB) *Database {
	return &Database{
		config: config,
	}
}

func (db *Database) Open() *pgx.Conn {
	config, err := pgx.ParseConnectionString(
		fmt.Sprintf(
			"user=%s password=%s dbname=%s sslmode=%s",
			db.config.Postgres.Username,
			db.config.Postgres.Password,
			db.config.Postgres.DbName,
			db.config.Postgres.SslMode,
		))
	if err != nil {
		Lg("config", "Database.Open").Error(err)
		return nil
	}

	database, err := pgx.Connect(config)
	if err != nil {
		Lg("config", "Database.Open").Error(err)
		return nil
	}

	return database
}
