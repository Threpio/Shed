package main

import (
	"context"
	"database/sql"
	_ "embed"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/threpio/shed/data"
	"log"
)

//go:embed schema.sql
var ddl string

func initDB() (*data.Queries, error) {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	queries := data.New(db)

	return queries, nil
}

func main() {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
		return
	}

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		log.Fatal(err)
		return
	}

	queries := data.New(db)

	ICompetition := NewICompetition(queries)

	r := gin.Default()
	r.GET("/ping", RHandlerPing())
	r.POST("/competition", RHandlerCompetitionCreate(ICompetition))
	r.GET("/competitions", RHandlerCompetitionList(ICompetition))
	r.Run() // listen and serve on 0.0.0.0:8080
}
