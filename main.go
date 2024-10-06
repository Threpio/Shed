package main

import (
	"context"
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/threpio/shed/data"
)

//go:embed schema.sql
var ddl string

func initDB() (*data.Queries, error) {

	ctx := context.Background()

	// Set up connection pool configuration (PostgreSQL)
	connString := "user=theoa dbname=postgres sslmode=verify-full"

	// Initialize the connection pool
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}

	queries := data.New(pool)

	return queries, nil
}

func main() {

	queries, err := initDB()
	if err != nil {
		panic(err)
	}

	ICompetition := NewICompetition(queries)
	ICompetitionDay := NewICompetitionDay(queries)
	IParentOrg := NewIParentOrg(queries)

	r := gin.Default()
	r.GET("/ping", RHandlerPing())
	// Competition Paths
	r.POST("/competitions", RHandlerCompetitionCreate(ICompetition))
	r.GET("/competitions", RHandlerCompetitionList(ICompetition))
	// Competition Day Paths
	r.POST("/competitions/days", RHandlerCompetitionDayCreate(ICompetitionDay))
	r.GET("/competitions/days", RHandlerCompetitionDayList(ICompetitionDay))

	r.POST("/parentorganisations", RHandlerParentOrgCreate(IParentOrg))
	r.GET("/parentorganisations", RHandlerParentOrgList(IParentOrg))

	r.Run() // listen and serve on 0.0.0.0:8080
}
