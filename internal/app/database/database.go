package database

import (
	"database/sql"
	"fmt"
	configDB "github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	_ "github.com/lib/pq"
)

type DB struct {
	dbPool *sql.DB
	config *configDB.ConfDB
}

func NewDB(config *configDB.ConfDB) *DB {
	return &DB{
		config: config,
	}
}

func (db *DB) Open() error {
	conn, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s sslmode=%s",
		db.config.Postgres.Username,
		db.config.Postgres.Password,
		db.config.Postgres.Host,
		db.config.Postgres.DbName,
		db.config.Postgres.SslMode,
	))
	if err != nil {
		return err
	}
	db.dbPool = conn
	return nil
}

func (db *DB) Close() {
	db.dbPool.Close()
}

func (db *DB) Begin() (*sql.Tx, error) {
	return db.dbPool.Begin()
}

func (db *DB) Exec(sql string, arguments ...interface{}) (sql.Result, error) {
	return db.dbPool.Exec(sql, arguments...)
}

func (db *DB) Query(sql string, optionsAndArgs ...interface{}) (*sql.Rows, error) {
	return db.dbPool.Query(sql, optionsAndArgs...)
}

func (db *DB) QueryRow(sql string, optionsAndArgs ...interface{}) *sql.Row {
	return db.dbPool.QueryRow(sql, optionsAndArgs...)
}
