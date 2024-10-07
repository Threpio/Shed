package main

import (
	"context"
	"fmt"
	"github.com/threpio/shed/data"
)

type HCompetitionDayCreateBody struct {
	CompetitionID int64
	Day           string
}

type HCompetitionDayCreateResponse struct {
	ID            int64
	CompetitionID int64
	Day           string
}

type HCompetitionDayListResponse struct {
	CompetitionDays []data.CompetitionDay
}

type ICompetitionDay struct {
	query *data.Queries
}

func NewICompetitionDay(query *data.Queries) *ICompetitionDay {
	return &ICompetitionDay{
		query: query,
	}
}

func (c *ICompetitionDay) Create(ctx context.Context, B HCompetitionDayCreateBody) (*HCompetitionDayCreateResponse, error) {
	insertedCompetitionDay, err := c.query.InsertCompetitionDay(ctx, data.InsertCompetitionDayParams{
		CompetitionID: B.CompetitionID,
		Day:           B.Day,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(insertedCompetitionDay)

	insertedCompetitionDayResponse := &HCompetitionDayCreateResponse{
		ID:            insertedCompetitionDay.ID,
		CompetitionID: insertedCompetitionDay.CompetitionID,
		Day:           insertedCompetitionDay.Day,
	}

	return insertedCompetitionDayResponse, nil
}

func (c *ICompetitionDay) List(ctx context.Context) (*HCompetitionDayListResponse, error) {
	print("List")
	list, err := c.query.ListCompetitionDays(ctx)
	if err != nil {
		return nil, err
	}

	listResponse := &HCompetitionDayListResponse{
		CompetitionDays: list,
	}

	return listResponse, nil
}
