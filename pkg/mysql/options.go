package mysql

import (
	"database/sql"
	"time"
)

type Options func(*sql.DB)

func SetConnMaxLifetime(d time.Duration) Options {
	return func(db *sql.DB) {
		db.SetConnMaxLifetime(d)
	}
}

func SetMaxIdleConns(n int) Options {
	return func(db *sql.DB) {
		db.SetMaxIdleConns(n)
	}
}

func SetMaxOpenConns(n int) Options {
	return func(db *sql.DB) {
		db.SetMaxOpenConns(n)
	}
}
