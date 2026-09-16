package database

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
)

type Connection struct {
	DB      *sql.DB
	Queries *Queries
}

func Connect(dbURI string) (*Connection, error) {
	dbConn, err := sql.Open("sqlite", dbURI)

	if err != nil {
		return nil, fmt.Errorf("Error opening db connection, %w", err)
	}

	if err := dbConn.Ping(); err != nil {
		dbConn.Close()
		return nil, fmt.Errorf("Error pinging db, %w", err)
	}

	queries := New(dbConn)

	return &Connection{
		DB:      dbConn,
		Queries: queries,
	}, nil
}
