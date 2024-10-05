package main

import (
	"context"
	"github.com/threpio/shed/data"
	"log"
	"time"
)

type AdminData struct {
	ctx     context.Context
	queries *data.Queries
}

func NewAdminData(ctx context.Context, queries *data.Queries) *AdminData {
	return &AdminData{
		ctx:     ctx,
		queries: queries,
	}
}

func (a *AdminData) CreateCompetition(name, instance string) (data.Competition, error) {
	log.Default().Println(time.Now(), " - Creating competition with name", name, "and instance", instance)
	return a.queries.InsertCompetition(a.ctx, data.InsertCompetitionParams{
		Name:     name,
		Instance: instance,
	})
}

func (a *AdminData) ListCompetitions() ([]data.Competition, error) {
	log.Default().Println(time.Now(), " - Listing all competitions")
	return a.queries.ListCompetitions(a.ctx)
}
