package main

import (
	"context"
	"database/sql"
	_ "embed"
	_ "github.com/mattn/go-sqlite3"
	"github.com/threpio/shed/data"
	"log"
)

//go:embed schema.sql
var ddl string

func run() error {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return err
	}

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}

	_ = data.New(db)

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

}
