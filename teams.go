package main

import (
	"context"
	"errors"
	"github.com/threpio/shed/data"
	"strings"
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
	switch strings.ToLower(B.PlayCategory) {
	case "mixed":
		break
	case "open":
		break
	case "womensPlus":
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

func (t *ITeam) List(ctx context.Context) (*HTeamListResponse, error) {
	//TODO: Decide whether the names can be duplicated
	list, err := t.query.ListTeams(ctx)
	if err != nil {
		return nil, err
	}

	listResponse := &HTeamListResponse{
		Teams: list,
	}

	return listResponse, nil
}
