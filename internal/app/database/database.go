package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Database struct {
	config   *Config
	database *sql.DB
}

func New(config *Config) *Database {
	return &Database{
		config: config,
	}
}

func (d *Database) Open() error {
	db, err := sql.Open(
		d.config.DriverName,
		fmt.Sprintf(
			"user=%s password=%s dbname=%s sslmode=%s",
			d.config.Username,
			d.config.Password,
			d.config.DbName,
			d.config.SslMode))
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
