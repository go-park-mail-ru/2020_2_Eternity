package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type IDbConn interface {
	Begin() (*sql.Tx, error)
	Exec(sql string, arguments ...interface{}) (sql.Result, error)
	Query(sql string, optionsAndArgs ...interface{}) (*sql.Rows, error)
	QueryRow(sql string, optionsAndArgs ...interface{}) *sql.Row
}
