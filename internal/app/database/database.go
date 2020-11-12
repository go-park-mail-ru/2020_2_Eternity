package database

import (
	"context"
	"fmt"
	configDB "github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	dbPool *pgxpool.Pool
	config *configDB.ConfDB
}

func NewDB(config *configDB.ConfDB) *DB {
	return &DB{
		config: config,
	}
}

func (db *DB) Open() error {
	conf, err := pgxpool.ParseConfig(fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s sslmode=%s pool_max_conns=%s",
		db.config.Postgres.Username,
		db.config.Postgres.Password,
		db.config.Postgres.Host,
		db.config.Postgres.DbName,
		db.config.Postgres.SslMode,
		db.config.Postgres.MaxConn,
	))
	if err != nil {
		return err
	}

	db.dbPool, err = pgxpool.ConnectConfig(context.Background(), conf)

	if err != nil {
		return err
	}
	return nil
}

func (db *DB) Close() {
	db.dbPool.Close()
}

func (db *DB) Begin(ctx context.Context) (pgx.Tx, error) {
	return db.dbPool.Begin(ctx)
}

func (db *DB) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return db.dbPool.Exec(ctx, sql, arguments...)
}

func (db *DB) Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error) {
	return db.dbPool.Query(ctx, sql, optionsAndArgs...)
}

func (db *DB) QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row {
	return db.dbPool.QueryRow(ctx, sql, optionsAndArgs...)
}
