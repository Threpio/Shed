package main

import (
	"context"
	"fmt"
	"github.com/threpio/shed/data"
)

type HCompetitionCreateBody struct {
	Name     string
	Instance string
}

type HCompetitionCreateResponse struct {
	ID       int64
	Name     string
	Instance string
}

type HCompetitionListResponse struct {
	Competitions []data.Competition
}

//
//type CompetitionInterface interface {
//	CreateCompetition(data.InsertCompetitionParams) (data.Competition, error)
//	ListCompetitions() ([]data.Competition, error)
//}

type ICompetition struct {
	query *data.Queries
}

func NewICompetition(query *data.Queries) *ICompetition {
	return &ICompetition{
		query: query,
	}
}

func (c *ICompetition) Create(ctx context.Context, B HCompetitionCreateBody) (*HCompetitionCreateResponse, error) {
	insertedCompetition, err := c.query.InsertCompetition(ctx, data.InsertCompetitionParams{
		Name:     B.Name,
		Instance: B.Instance,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(insertedCompetition)

	insertedCompetitionResponse := &HCompetitionCreateResponse{
		ID:       insertedCompetition.ID,
		Name:     insertedCompetition.Name,
		Instance: insertedCompetition.Instance,
	}

	return insertedCompetitionResponse, nil
}

func (c *ICompetition) List(ctx context.Context) (*HCompetitionListResponse, error) {
	print("List")
	competitions, err := c.query.ListCompetitions(ctx)
	if err != nil {
		return nil, err
	}
	return &HCompetitionListResponse{
		Competitions: competitions,
	}, nil
}
