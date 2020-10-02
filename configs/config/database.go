package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Database struct {
	config   *ConfDB
	database *sql.DB
}

func newDatabase(config *ConfDB) *Database {
	return &Database{
		config: config,
	}
}

func (d *Database) Open() error {
	db, err := sql.Open(
		d.config.Postgres.DriverName,
		fmt.Sprintf(
			"user=%s password=%s dbname=%s sslmode=%s",
			d.config.Postgres.Username,
			d.config.Postgres.Password,
			d.config.Postgres.DbName,
			d.config.Postgres.SslMode))
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return err
	}

	d.database = db
	return nil
}

func (d *Database) Close() error {
	return d.database.Close()
}
