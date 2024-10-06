package main

import "github.com/threpio/shed/data"

type HCompetitionDayCreateBody struct {
	CompetitionID int64  `json:"competition_id"`
	Day           string `json:"day"`
}

type HCompetitionDayCreateResponse struct {
	ID            int64  `json:"id"`
	CompetitionID int64  `json:"competition_id"`
	Day           string `json:"day"`
}

type HCompetitionDayListResponse struct {
	CompetitionDays []data.CompetitionDay `json:"competition_days"`
}

type ICompetitionDay struct {
	query *data.Queries
}

func NewICompetitionDay(query *data.Queries) *ICompetitionDay {
	return &ICompetitionDay{
		query: query,
	}
}
