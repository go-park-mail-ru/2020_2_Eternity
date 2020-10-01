package database

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs"
	"github.com/jackc/pgx"
)

type DB struct {
	conn *pgx.Conn
}

func NewConnection() *DB {
	return &DB{}
}

func (db *DB) Open(c *configs.DBConfig) error {
	config := pgx.ConnConfig{
		Host:     c.Host,
		Port:     c.Port,
		Database: c.Database,
		User:     c.User,
		Password: c.Password,
	}
	var err error
	db.conn, err = pgx.Connect(config)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) Close() error {
	if err := db.conn.Close(); err != nil {
		return err
	}
	return nil
}

func InitDB(c *configs.DBConfig) (*DB, error) {
	conn := NewConnection()
	if err := conn.Open(c); err != nil {
		return nil, err
	}
	return conn, nil
}
