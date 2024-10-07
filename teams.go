package main

import (
	"context"
	"errors"
	"github.com/threpio/shed/data"
)

type HTeamCreateBody struct {
	Name                 string
	CompetitionID        int64
	ParentOrganisationID int64
	PlayCategory         string
}

type HTeamCreateResponse struct {
	ID                   int64
	Name                 string
	CompetitionID        int64
	ParentOrganisationID int64
	PlayCategory         string
}

type HTeamListResponse struct {
	Teams []data.Team
}

type ITeam struct {
	query *data.Queries
}

func NewITeam(query *data.Queries) *ITeam {
	return &ITeam{
		query: query,
	}
}

func (t *ITeam) Create(ctx context.Context, B HTeamCreateBody) (*HTeamCreateResponse, error) {
	// TODO: Decide if we want to use the InsertTeamParameters or generate a new struct

	//TODO: Check that the play category is specific from the list

	// 'Mixed','Open','WomensPlus'
	switch B.PlayCategory {
	case "Mixed":
		break
	case "Open":
		break
	case "WomensPlus":
		break
	default:
		return nil, errors.New("Play Category not valid")
	}

	i := data.InsertTeamParams{
		Name:                 B.Name,
		CompetitionID:        B.CompetitionID,
		ParentOrganisationID: B.ParentOrganisationID,
		PlayCategory:         B.PlayCategory,
	}

	insertedTeam, err := t.query.InsertTeam(ctx, i)
	if err != nil {
		return nil, err
	}

	insertedTeamResponse := &HTeamCreateResponse{
		ID:                   insertedTeam.ID,
		Name:                 insertedTeam.Name,
		CompetitionID:        insertedTeam.CompetitionID,
		ParentOrganisationID: insertedTeam.ParentOrganisationID,
		PlayCategory:         insertedTeam.PlayCategory.(string),
	}

	return insertedTeamResponse, nil
}
