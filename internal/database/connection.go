package database

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
	"github.com/subratamondal1029/goTodo/internal/cli"
)

type Connection struct {
	db      *sql.DB
	queries *Queries
}

func Connect() (*Connection, error) {
	dbConn, err := sql.Open("sqlite", cli.DBURI)

	if err != nil {
		return nil, fmt.Errorf("Error opening db connection, %w", err)
	}
	defer dbConn.Close()

	queries := New(dbConn)

	return &Connection{
		db:      dbConn,
		queries: queries,
	}, nil
}
