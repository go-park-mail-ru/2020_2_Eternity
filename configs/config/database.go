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

func (d *Database) Open() error {
	config, err := pgx.ParseConnectionString(
		fmt.Sprintf(
			"user=%s password=%s dbname=%s sslmode=%s",
			d.config.Postgres.Username,
			d.config.Postgres.Password,
			d.config.Postgres.DbName,
			d.config.Postgres.SslMode,
			))
	if err != nil {
		return err
	}

	d.database, err = pgx.Connect(config)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) Close() error {
	return d.database.Close()
}
