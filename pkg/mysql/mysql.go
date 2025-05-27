package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	DB *sql.DB
}

func New(url string, opts ...Options) (*MySQL, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, fmt.Errorf("pkg/mysql - New - sql.Open: %w", err)
	}

	err = db.Ping() // Need to do this to check that the connection is valid
	if err != nil {
		return nil, fmt.Errorf("pkg/mysql - New - db.Ping: %w", err)
	}

	for _, opt := range opts {
		opt(db)
	}
	return &MySQL{DB: db}, nil
}

func (m *MySQL) Close() error {
	return m.DB.Close()
}
